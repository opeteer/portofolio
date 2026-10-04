package bench

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type SysInfo struct {
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	NumCPU    int    `json:"num_cpu"`
	GoVersion string `json:"go_version"`
	Compiler  string `json:"compiler"`
	MaxProcs  int    `json:"gomaxprocs"`
}

func GetSysInfo() SysInfo {
	return SysInfo{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		NumCPU:    runtime.NumCPU(),
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		MaxProcs:  runtime.GOMAXPROCS(0),
	}
}

type CPUBenchResult struct {
	Duration       time.Duration `json:"duration"`
	DurationMs     float64       `json:"duration_ms"`
	Workers        int           `json:"workers"`
	TotalOps       uint64        `json:"total_ops"`
	OpsPerSec      float64       `json:"ops_per_sec"`
	BytesProcessed uint64        `json:"bytes_processed"`
	ThroughputMBs  float64       `json:"throughput_mb_s"`
	ThroughputGBs  float64       `json:"throughput_gb_s"`
}

type MemBenchResult struct {
	Duration       time.Duration `json:"duration"`
	DurationMs     float64       `json:"duration_ms"`
	Workers        int           `json:"workers"`
	TotalAllocMB   float64       `json:"total_alloc_mb"`
	AllocRateMBs   float64       `json:"alloc_rate_mb_s"`
	MallocsCount   uint64        `json:"mallocs_count"`
	MallocsPerSec  float64       `json:"mallocs_per_sec"`
	NumGCCycles    uint32        `json:"num_gc_cycles"`
	GCPauseTotalMs float64       `json:"gc_pause_total_ms"`
}

type LatencyDistribution struct {
	MinMs float64 `json:"min_ms"`
	AvgMs float64 `json:"avg_ms"`
	P50Ms float64 `json:"p50_ms"`
	P90Ms float64 `json:"p90_ms"`
	P99Ms float64 `json:"p99_ms"`
	MaxMs float64 `json:"max_ms"`
}

type EngineBenchResult struct {
	Duration        time.Duration       `json:"duration"`
	DurationMs      float64             `json:"duration_ms"`
	Workers         int                 `json:"workers"`
	TotalRequests   uint64              `json:"total_requests"`
	RPS             float64             `json:"rps"`
	Latency         LatencyDistribution `json:"latency"`
	EndpointsTested []string            `json:"endpoints_tested"`
}

type FullReport struct {
	Timestamp   time.Time          `json:"timestamp"`
	SysInfo     SysInfo            `json:"system_info"`
	Target      string             `json:"target"`
	ScoreRating string             `json:"score_rating"`
	CPU         *CPUBenchResult    `json:"cpu,omitempty"`
	Memory      *MemBenchResult    `json:"memory,omitempty"`
	Engine      *EngineBenchResult `json:"engine,omitempty"`
}

type Options struct {
	Duration    time.Duration
	Concurrency int
	Target      string
	Handler     http.Handler
}

// RunCPUBenchmark performs parallel SHA-256 cryptographic throughput testing
func RunCPUBenchmark(concurrency int, duration time.Duration) CPUBenchResult {
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}
	if duration <= 0 {
		duration = 1 * time.Second
	}

	blockSize := 64 * 1024 // 64 KB per block
	chunk := make([]byte, blockSize)
	for i := range chunk {
		chunk[i] = byte(i % 256)
	}

	var ops atomic.Uint64
	var totalBytes atomic.Uint64
	stop := make(chan struct{})

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h := sha256.New()
			for {
				select {
				case <-stop:
					return
				default:
					h.Reset()
					h.Write(chunk)
					h.Sum(nil)
					ops.Add(1)
					totalBytes.Add(uint64(blockSize))
				}
			}
		}()
	}

	time.Sleep(duration)
	close(stop)
	wg.Wait()
	elapsed := time.Since(start)

	totalOps := ops.Load()
	mbProcessed := float64(totalBytes.Load()) / (1024 * 1024)
	throughputMBs := mbProcessed / elapsed.Seconds()
	opsPerSec := float64(totalOps) / elapsed.Seconds()

	return CPUBenchResult{
		Duration:       elapsed,
		DurationMs:     float64(elapsed.Microseconds()) / 1000.0,
		Workers:        concurrency,
		TotalOps:       totalOps,
		OpsPerSec:      opsPerSec,
		BytesProcessed: totalBytes.Load(),
		ThroughputMBs:  throughputMBs,
		ThroughputGBs:  throughputMBs / 1024.0,
	}
}

// RunMemoryBenchmark performs concurrent allocation and measures GC metrics
func RunMemoryBenchmark(concurrency int, duration time.Duration) MemBenchResult {
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}
	if duration <= 0 {
		duration = 1 * time.Second
	}

	runtime.GC()
	var mBefore runtime.MemStats
	runtime.ReadMemStats(&mBefore)

	stop := make(chan struct{})
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			var holder [][]byte
			counter := 0
			for {
				select {
				case <-stop:
					_ = holder
					return
				default:
					// Allocate varying chunk sizes (1KB to 32KB)
					size := 1024 * ((counter % 32) + 1)
					buf := make([]byte, size)
					buf[0] = byte(workerID)
					holder = append(holder, buf)
					counter++

					if len(holder) > 1000 {
						holder = nil // Release to trigger GC
					}
				}
			}
		}(i)
	}

	time.Sleep(duration)
	close(stop)
	wg.Wait()
	elapsed := time.Since(start)

	var mAfter runtime.MemStats
	runtime.ReadMemStats(&mAfter)

	bytesAllocated := mAfter.TotalAlloc - mBefore.TotalAlloc
	mallocs := mAfter.Mallocs - mBefore.Mallocs
	gcCycles := mAfter.NumGC - mBefore.NumGC
	gcPauseNs := mAfter.PauseTotalNs - mBefore.PauseTotalNs

	totalAllocMB := float64(bytesAllocated) / (1024 * 1024)
	allocRateMBs := totalAllocMB / elapsed.Seconds()
	mallocsPerSec := float64(mallocs) / elapsed.Seconds()
	gcPauseMs := float64(gcPauseNs) / 1e6

	return MemBenchResult{
		Duration:       elapsed,
		DurationMs:     float64(elapsed.Microseconds()) / 1000.0,
		Workers:        concurrency,
		TotalAllocMB:   totalAllocMB,
		AllocRateMBs:   allocRateMBs,
		MallocsCount:   mallocs,
		MallocsPerSec:  mallocsPerSec,
		NumGCCycles:    gcCycles,
		GCPauseTotalMs: gcPauseMs,
	}
}

// fallbackHandler provides a fast mock router if no HTTP handler is provided
func fallbackHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "ok",
			"bench":  true,
		})
	})
	mux.HandleFunc("/api/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"username": "opeteer",
			"role":     "systems_engineer",
		})
	})
	return mux
}

// RunEngineBenchmark benchmarks internal HTTP router and JSON serialization
func RunEngineBenchmark(concurrency int, duration time.Duration, handler http.Handler) EngineBenchResult {
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}
	if duration <= 0 {
		duration = 1 * time.Second
	}
	if handler == nil {
		handler = fallbackHandler()
	}

	endpoints := []string{"/api/profile", "/api/projects", "/api/stats"}
	var totalReqs atomic.Uint64

	// Sample latencies for percentile calculations (up to 5,000 samples to keep memory bound)
	var latenciesMu sync.Mutex
	var sampleLatencies []float64
	maxSamples := 5000

	stop := make(chan struct{})
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(wID int) {
			defer wg.Done()
			reqIndex := wID
			var localSamples []float64

			for {
				select {
				case <-stop:
					if len(localSamples) > 0 {
						latenciesMu.Lock()
						if len(sampleLatencies) < maxSamples {
							remaining := maxSamples - len(sampleLatencies)
							if len(localSamples) > remaining {
								sampleLatencies = append(sampleLatencies, localSamples[:remaining]...)
							} else {
								sampleLatencies = append(sampleLatencies, localSamples...)
							}
						}
						latenciesMu.Unlock()
					}
					return
				default:
					path := endpoints[reqIndex%len(endpoints)]
					req := httptest.NewRequest("GET", path, nil)
					w := httptest.NewRecorder()

					t0 := time.Now()
					handler.ServeHTTP(w, req)
					latMs := float64(time.Since(t0).Microseconds()) / 1000.0

					totalReqs.Add(1)
					if len(localSamples) < (maxSamples/concurrency)+10 {
						localSamples = append(localSamples, latMs)
					}
					reqIndex++
				}
			}
		}(i)
	}

	time.Sleep(duration)
	close(stop)
	wg.Wait()
	elapsed := time.Since(start)

	completedReqs := totalReqs.Load()
	rps := float64(completedReqs) / elapsed.Seconds()

	// Compute latency distribution
	dist := LatencyDistribution{}
	if len(sampleLatencies) > 0 {
		sort.Float64s(sampleLatencies)
		dist.MinMs = sampleLatencies[0]
		dist.MaxMs = sampleLatencies[len(sampleLatencies)-1]

		var sum float64
		for _, v := range sampleLatencies {
			sum += v
		}
		dist.AvgMs = sum / float64(len(sampleLatencies))

		p50Idx := int(float64(len(sampleLatencies)) * 0.50)
		p90Idx := int(float64(len(sampleLatencies)) * 0.90)
		p99Idx := int(float64(len(sampleLatencies)) * 0.99)

		if p50Idx >= len(sampleLatencies) {
			p50Idx = len(sampleLatencies) - 1
		}
		if p90Idx >= len(sampleLatencies) {
			p90Idx = len(sampleLatencies) - 1
		}
		if p99Idx >= len(sampleLatencies) {
			p99Idx = len(sampleLatencies) - 1
		}

		dist.P50Ms = sampleLatencies[p50Idx]
		dist.P90Ms = sampleLatencies[p90Idx]
		dist.P99Ms = sampleLatencies[p99Idx]
	}

	return EngineBenchResult{
		Duration:        elapsed,
		DurationMs:      float64(elapsed.Microseconds()) / 1000.0,
		Workers:         concurrency,
		TotalRequests:   completedReqs,
		RPS:             rps,
		Latency:         dist,
		EndpointsTested: endpoints,
	}
}

// RunSuite runs the requested benchmark suite
func RunSuite(opts Options) FullReport {
	if opts.Duration <= 0 {
		opts.Duration = 1 * time.Second
	}
	if opts.Concurrency <= 0 {
		opts.Concurrency = runtime.NumCPU()
	}
	if opts.Target == "" {
		opts.Target = "all"
	}

	report := FullReport{
		Timestamp: time.Now(),
		SysInfo:   GetSysInfo(),
		Target:    opts.Target,
	}

	switch opts.Target {
	case "cpu":
		res := RunCPUBenchmark(opts.Concurrency, opts.Duration)
		report.CPU = &res
	case "mem", "memory":
		res := RunMemoryBenchmark(opts.Concurrency, opts.Duration)
		report.Memory = &res
	case "engine", "http", "router":
		res := RunEngineBenchmark(opts.Concurrency, opts.Duration, opts.Handler)
		report.Engine = &res
	default:
		// all
		cRes := RunCPUBenchmark(opts.Concurrency, opts.Duration)
		report.CPU = &cRes

		mRes := RunMemoryBenchmark(opts.Concurrency, opts.Duration)
		report.Memory = &mRes

		eRes := RunEngineBenchmark(opts.Concurrency, opts.Duration, opts.Handler)
		report.Engine = &eRes
	}

	// Calculate Score Rating
	report.ScoreRating = calculateScoreRating(report)
	return report
}

func calculateScoreRating(r FullReport) string {
	score := 0
	maxPossible := 0

	if r.CPU != nil {
		maxPossible += 3
		if r.CPU.ThroughputGBs >= 8.0 {
			score += 3
		} else if r.CPU.ThroughputGBs >= 2.0 {
			score += 2
		} else {
			score += 1
		}
	}
	if r.Memory != nil {
		maxPossible += 3
		if r.Memory.AllocRateMBs >= 1000.0 {
			score += 3
		} else if r.Memory.AllocRateMBs >= 300.0 {
			score += 2
		} else {
			score += 1
		}
	}
	if r.Engine != nil {
		maxPossible += 3
		if r.Engine.RPS >= 40000.0 {
			score += 3
		} else if r.Engine.RPS >= 10000.0 {
			score += 2
		} else {
			score += 1
		}
	}

	if maxPossible == 0 {
		return "N/A"
	}

	ratio := float64(score) / float64(maxPossible)
	switch {
	case ratio >= 0.8:
		return "Ultra Tier (Elite Multi-Core Workstation / Cloud Core)"
	case ratio >= 0.5:
		return "High Tier (Production Cloud Node / Fast Developer Machine)"
	default:
		return "Standard Tier (Efficient Core / Micro Cloud Node)"
	}
}

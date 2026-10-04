package bench

import (
	"net/http"
	"testing"
	"time"
)

func TestCPUBenchmark(t *testing.T) {
	res := RunCPUBenchmark(2, 50*time.Millisecond)
	if res.TotalOps == 0 {
		t.Errorf("expected TotalOps > 0, got %d", res.TotalOps)
	}
	if res.OpsPerSec <= 0 {
		t.Errorf("expected OpsPerSec > 0, got %f", res.OpsPerSec)
	}
	if res.ThroughputMBs <= 0 {
		t.Errorf("expected ThroughputMBs > 0, got %f", res.ThroughputMBs)
	}
	if res.Workers != 2 {
		t.Errorf("expected 2 workers, got %d", res.Workers)
	}
}

func TestMemoryBenchmark(t *testing.T) {
	res := RunMemoryBenchmark(2, 50*time.Millisecond)
	if res.MallocsCount == 0 {
		t.Errorf("expected MallocsCount > 0, got %d", res.MallocsCount)
	}
	if res.AllocRateMBs <= 0 {
		t.Errorf("expected AllocRateMBs > 0, got %f", res.AllocRateMBs)
	}
	if res.Workers != 2 {
		t.Errorf("expected 2 workers, got %d", res.Workers)
	}
}

func TestEngineBenchmark(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/api/profile", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"username":"opeteer"}`))
	})
	mux.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	})

	res := RunEngineBenchmark(2, 50*time.Millisecond, mux)
	if res.TotalRequests == 0 {
		t.Errorf("expected TotalRequests > 0, got %d", res.TotalRequests)
	}
	if res.RPS <= 0 {
		t.Errorf("expected RPS > 0, got %f", res.RPS)
	}
	if res.Latency.AvgMs < 0 {
		t.Errorf("invalid avg latency: %f", res.Latency.AvgMs)
	}
}

func TestRunSuite(t *testing.T) {
	opts := Options{
		Duration:    50 * time.Millisecond,
		Concurrency: 2,
		Target:      "all",
	}
	report := RunSuite(opts)

	if report.SysInfo.NumCPU <= 0 {
		t.Errorf("expected NumCPU > 0, got %d", report.SysInfo.NumCPU)
	}
	if report.CPU == nil {
		t.Errorf("expected CPU bench result in report")
	}
	if report.Memory == nil {
		t.Errorf("expected Memory bench result in report")
	}
	if report.Engine == nil {
		t.Errorf("expected Engine bench result in report")
	}
	if report.ScoreRating == "" {
		t.Errorf("expected non-empty score rating")
	}
}

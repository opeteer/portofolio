package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/opeteer/opeteer-cli/pkg/bench"
	"github.com/opeteer/opeteer-cli/pkg/server"
	"github.com/opeteer/opeteer-cli/pkg/ui"
)

func PrintBenchHelp() {
	ui.PrintSectionHeader("OPETEER CLI: bench command")
	fmt.Println("USAGE:")
	fmt.Println("  opeteer bench [flags]")
	fmt.Println()
	fmt.Println("ALIASES:")
	fmt.Println("  benchmark, speed")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  Executes a real-time (100% non-mock) systems micro-benchmark suite measuring")
	fmt.Println("  CPU cryptographic hashing throughput, Go runtime memory allocator/GC stress,")
	fmt.Println("  and embedded HTTP router & JSON serializer performance.")
	fmt.Println()
	fmt.Println("FLAGS:")
	fmt.Println("  -d, --duration <time>     Duration per benchmark (default: 1s, e.g. 500ms, 2s)")
	fmt.Println("  -c, --concurrency <num>   Number of concurrent goroutines (default: NumCPU)")
	fmt.Println("  -t, --target <suite>      Benchmark target: all, cpu, mem, engine (default: all)")
	fmt.Println("  -j, --json                Output benchmark results in raw JSON format")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  $ opeteer bench")
	fmt.Println("  $ opeteer bench --target cpu --duration 2s")
	fmt.Println("  $ opeteer bench --target engine -c 8")
	fmt.Println("  $ opeteer bench --json | jq .score_rating")
	fmt.Println()
}

func ParseBenchFlags(args []string) (bench.Options, bool) {
	opts := bench.Options{
		Duration:    1 * time.Second,
		Concurrency: runtime.NumCPU(),
		Target:      "all",
	}
	jsonOut := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "-j" || arg == "--json":
			jsonOut = true
		case arg == "-d" || arg == "--duration":
			if i+1 < len(args) {
				if d, err := time.ParseDuration(args[i+1]); err == nil && d > 0 {
					opts.Duration = d
				}
				i++
			}
		case strings.HasPrefix(arg, "--duration="):
			if d, err := time.ParseDuration(strings.TrimPrefix(arg, "--duration=")); err == nil && d > 0 {
				opts.Duration = d
			}
		case arg == "-c" || arg == "--concurrency":
			if i+1 < len(args) {
				if n, err := strconv.Atoi(args[i+1]); err == nil && n > 0 {
					opts.Concurrency = n
				}
				i++
			}
		case strings.HasPrefix(arg, "--concurrency="):
			if n, err := strconv.Atoi(strings.TrimPrefix(arg, "--concurrency=")); err == nil && n > 0 {
				opts.Concurrency = n
			}
		case arg == "-t" || arg == "--target":
			if i+1 < len(args) {
				opts.Target = strings.ToLower(args[i+1])
				i++
			}
		case strings.HasPrefix(arg, "--target="):
			opts.Target = strings.ToLower(strings.TrimPrefix(arg, "--target="))
		}
	}

	return opts, jsonOut
}

func RunBench(args []string) {
	opts, jsonOut := ParseBenchFlags(args)

	// Attach server mux for engine benchmark (without request logging)
	opts.Handler = server.NewRawMux()

	if jsonOut {
		report := bench.RunSuite(opts)
		data, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			ui.Fatal("Failed to encode benchmark report to JSON: %v", err)
		}
		fmt.Println(string(data))
		return
	}

	ui.PrintSectionHeader("OPETEER SYSTEMS BENCHMARK SUITE")

	sys := bench.GetSysInfo()
	fmt.Println(ui.Cyan("┌─ Host System & Go Runtime Diagnostics"))
	fmt.Printf("│  %-18s: %s\n", "OS / Architecture", ui.Bold(fmt.Sprintf("%s / %s", sys.OS, sys.Arch)))
	fmt.Printf("│  %-18s: %s\n", "Logical CPU Cores", ui.Bold(fmt.Sprintf("%d cores (GOMAXPROCS: %d)", sys.NumCPU, sys.MaxProcs)))
	fmt.Printf("│  %-18s: %s\n", "Go Compiler / Ver", ui.White(fmt.Sprintf("%s (%s)", sys.GoVersion, sys.Compiler)))
	fmt.Printf("│  %-18s: %s\n", "Bench Concurrency", ui.BrightCyan(fmt.Sprintf("%d parallel workers", opts.Concurrency)))
	fmt.Printf("│  %-18s: %s\n", "Sample Duration", ui.BrightCyan(opts.Duration.String()))
	fmt.Println(ui.Cyan("└──────────────────────────────────────────────────"))
	fmt.Println()

	fmt.Print(ui.Dim("Executing benchmark suite... "))

	report := bench.RunSuite(opts)
	fmt.Println(ui.Green("Done!"))
	fmt.Println()

	// 1. CPU Benchmark Result
	if report.CPU != nil {
		fmt.Println(ui.Yellow("┌─ [1/3] CPU Cryptographic Throughput (SHA-256 Multi-Core)"))
		fmt.Printf("│  %-18s: %s\n", "Hash Operations", ui.Bold(fmt.Sprintf("%d hashes", report.CPU.TotalOps)))
		fmt.Printf("│  %-18s: %s\n", "Hash Frequency", ui.BrightCyan(fmt.Sprintf("%.0f ops/sec", report.CPU.OpsPerSec)))
		fmt.Printf("│  %-18s: %s\n", "Data Processed", ui.White(fmt.Sprintf("%.2f MB in %v", float64(report.CPU.BytesProcessed)/(1024*1024), report.CPU.Duration.Round(time.Millisecond))))
		fmt.Printf("│  %-18s: %s\n", "Crypto Throughput", ui.Green(fmt.Sprintf("%.2f MB/s  (%.2f GB/s)", report.CPU.ThroughputMBs, report.CPU.ThroughputGBs)))
		fmt.Println(ui.Yellow("└──────────────────────────────────────────────────"))
		fmt.Println()
	}

	// 2. Memory Benchmark Result
	if report.Memory != nil {
		fmt.Println(ui.Magenta("┌─ [2/3] Memory Allocator & GC Pressure"))
		fmt.Printf("│  %-18s: %s\n", "Heap Allocated", ui.Bold(fmt.Sprintf("%.2f MB", report.Memory.TotalAllocMB)))
		fmt.Printf("│  %-18s: %s\n", "Alloc Rate", ui.BrightCyan(fmt.Sprintf("%.2f MB/s", report.Memory.AllocRateMBs)))
		fmt.Printf("│  %-18s: %s\n", "Object Mallocs", ui.White(fmt.Sprintf("%d objects (%.0f allocs/sec)", report.Memory.MallocsCount, report.Memory.MallocsPerSec)))
		fmt.Printf("│  %-18s: %s\n", "GC Execution", ui.Green(fmt.Sprintf("%d cycles completed (Pause: %.3f ms)", report.Memory.NumGCCycles, report.Memory.GCPauseTotalMs)))
		fmt.Println(ui.Magenta("└──────────────────────────────────────────────────"))
		fmt.Println()
	}

	// 3. HTTP & JSON Engine Result
	if report.Engine != nil {
		fmt.Println(ui.Cyan("┌─ [3/3] Embedded HTTP Router & JSON Serializer Engine"))
		fmt.Printf("│  %-18s: %s\n", "Total Requests", ui.Bold(fmt.Sprintf("%d requests", report.Engine.TotalRequests)))
		fmt.Printf("│  %-18s: %s\n", "Throughput (RPS)", ui.Green(fmt.Sprintf("%.0f req/sec", report.Engine.RPS)))
		fmt.Printf("│  %-18s: %s\n", "Avg Latency", ui.White(fmt.Sprintf("%.4f ms", report.Engine.Latency.AvgMs)))
		fmt.Printf("│  %-18s: %s\n", "Latency P50 / P90", ui.BrightCyan(fmt.Sprintf("%.4f ms / %.4f ms", report.Engine.Latency.P50Ms, report.Engine.Latency.P90Ms)))
		fmt.Printf("│  %-18s: %s\n", "Latency P99 / Max", ui.Yellow(fmt.Sprintf("%.4f ms / %.4f ms", report.Engine.Latency.P99Ms, report.Engine.Latency.MaxMs)))
		fmt.Println(ui.Cyan("└──────────────────────────────────────────────────"))
		fmt.Println()
	}

	// Summary Performance Index
	fmt.Printf("  %s %s\n", ui.Bold("System Performance Index:"), ui.BrightGreen(report.ScoreRating))
	fmt.Printf("  %s %s\n", ui.Dim("Tip:"), ui.Dim("Run 'opeteer bench --target cpu --duration 2s' for deep cryptographic stress."))
	fmt.Println()
}

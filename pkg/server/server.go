package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/opeteer/opeteer-cli/pkg/bench"
	"github.com/opeteer/opeteer-cli/pkg/data"
	"github.com/opeteer/opeteer-cli/pkg/live"
	"github.com/opeteer/opeteer-cli/pkg/ui"
	"github.com/opeteer/opeteer-cli/web"
)

type Config struct {
	Host string
	Port int
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		statusStr := fmt.Sprintf("%d", rec.statusCode)
		switch {
		case rec.statusCode >= 200 && rec.statusCode < 300:
			statusStr = ui.Green(statusStr)
		case rec.statusCode >= 300 && rec.statusCode < 400:
			statusStr = ui.Cyan(statusStr)
		case rec.statusCode >= 400 && rec.statusCode < 500:
			statusStr = ui.Yellow(statusStr)
		default:
			statusStr = ui.Red(statusStr)
		}

		timestamp := time.Now().Format("15:04:05")
		fmt.Printf("[%s] %s %-6s %-24s %s (%s)\n",
			ui.Dim(timestamp),
			statusStr,
			ui.Bold(r.Method),
			r.URL.Path,
			ui.Dim(duration.String()),
			r.RemoteAddr,
		)
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func NewMux() http.Handler {
	return loggingMiddleware(NewRawMux())
}

func NewRawMux() *http.ServeMux {
	mux := http.NewServeMux()

	// 1. Root and Static Web Assets
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/index.html" {
			http.NotFound(w, r)
			return
		}
		htmlData, err := web.Files.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Failed to load index.html", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(htmlData)
	})

	// 2. REST API: Profile
	mux.HandleFunc("/api/profile", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, data.MyProfile)
	})

	// 3. REST API: Projects (with category/lang filter)
	mux.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		qCat := strings.ToLower(r.URL.Query().Get("category"))
		qLang := strings.ToLower(r.URL.Query().Get("lang"))

		var result []data.Project
		for _, p := range data.AllProjects {
			if qCat != "" && !strings.Contains(strings.ToLower(p.Category), qCat) {
				continue
			}
			if qLang != "" && !strings.Contains(strings.ToLower(p.Language), qLang) {
				continue
			}
			result = append(result, p)
		}
		writeJSON(w, http.StatusOK, result)
	})

	// 4. REST API: Flagship Projects
	mux.HandleFunc("/api/projects/flagship", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, data.GetFlagshipProjects())
	})

	// 5. REST API: Skills
	mux.HandleFunc("/api/skills", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, data.TechnicalSkills)
	})

	// 6. REST API: Stats
	mux.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		langs := make(map[string]int)
		cats := make(map[string]int)
		totalSize := 0
		flagshipCount := 0

		for _, p := range data.AllProjects {
			totalSize += p.SizeKB
			if p.Flagship {
				flagshipCount++
			}
			lang := p.Language
			if lang == "" {
				lang = "Other"
			}
			langs[lang]++
			cats[p.Category]++
		}

		stats := map[string]interface{}{
			"total_repos":      len(data.AllProjects),
			"flagship_count":   flagshipCount,
			"total_size_kb":    totalSize,
			"languages":        langs,
			"categories":       cats,
			"primary_language": "Go / Python / Cloud",
		}
		writeJSON(w, http.StatusOK, stats)
	})

	// 7. REST API: Showtime (ASCII Art)
	mux.HandleFunc("/api/showtime", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"braille": ui.LogoBraille,
			"block":   ui.LogoBlock,
		})
	})

	// 8. REST API: Live Telemetry
	mux.HandleFunc("/api/live", func(w http.ResponseWriter, r *http.Request) {
		limit := 5
		if qLimit := r.URL.Query().Get("limit"); qLimit != "" {
			if parsed, err := strconv.Atoi(qLimit); err == nil && parsed > 0 {
				limit = parsed
			}
		}
		client := live.NewClient(3 * time.Second)
		report, _ := client.FetchLiveTelemetry(limit)
		writeJSON(w, http.StatusOK, report)
	})

	// 9. REST API: Systems Benchmark
	mux.HandleFunc("/api/bench", func(w http.ResponseWriter, r *http.Request) {
		duration := 300 * time.Millisecond
		if qDur := r.URL.Query().Get("duration"); qDur != "" {
			if d, err := time.ParseDuration(qDur); err == nil && d > 0 {
				duration = d
			}
		}
		concurrency := runtime.NumCPU()
		if qConc := r.URL.Query().Get("concurrency"); qConc != "" {
			if n, err := strconv.Atoi(qConc); err == nil && n > 0 {
				concurrency = n
			}
		}
		target := r.URL.Query().Get("target")
		if target == "" {
			target = "all"
		}
		report := bench.RunSuite(bench.Options{
			Duration:    duration,
			Concurrency: concurrency,
			Target:      target,
			Handler:     mux,
		})
		writeJSON(w, http.StatusOK, report)
	})

	// 10. Healthcheck Endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"version":   "1.0.0",
		})
	})

	return mux
}

func StartServer(cfg Config) error {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: NewMux(),
	}

	ui.PrintSectionHeader("OPETEER CLOUD PORTFOLIO SERVER")
	fmt.Printf("  %-18s: %s\n", ui.Bold("Web Portfolio"), ui.BrightCyan(fmt.Sprintf("http://localhost:%d", cfg.Port)))
	fmt.Printf("  %-18s: %s\n", ui.Bold("Profile API"), ui.White(fmt.Sprintf("http://localhost:%d/api/profile", cfg.Port)))
	fmt.Printf("  %-18s: %s\n", ui.Bold("Projects API"), ui.White(fmt.Sprintf("http://localhost:%d/api/projects", cfg.Port)))
	fmt.Printf("  %-18s: %s\n", ui.Bold("Emblem API"), ui.White(fmt.Sprintf("http://localhost:%d/api/showtime", cfg.Port)))
	fmt.Printf("  %-18s: %s\n", ui.Bold("Live API"), ui.BrightCyan(fmt.Sprintf("http://localhost:%d/api/live", cfg.Port)))
	fmt.Printf("  %-18s: %s\n", ui.Bold("Bench API"), ui.BrightCyan(fmt.Sprintf("http://localhost:%d/api/bench", cfg.Port)))
	fmt.Printf("  %-18s: %s\n", ui.Bold("Health Probe"), ui.Green(fmt.Sprintf("http://localhost:%d/healthz", cfg.Port)))
	fmt.Println()
	fmt.Println(ui.Dim("  Press Ctrl+C to terminate the web server."))
	fmt.Println(ui.Gray("  " + strings.Repeat("─", 65)))
	fmt.Println()

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		fmt.Println("\n" + ui.Yellow("Shutting down portfolio server gracefully..."))
		server.Close()
	}()

	return server.ListenAndServe()
}

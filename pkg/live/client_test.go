package live

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLiveTelemetrySuccess(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Limit", "60")
		w.Header().Set("X-RateLimit-Remaining", "55")
		w.Header().Set("X-RateLimit-Reset", "1700000000")

		if r.URL.Path == "/users/opeteer" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"username":     "opeteer",
				"name":         "Gerardo M Ardianta",
				"public_repos": 31,
				"followers":    4,
				"following":    12,
			})
			return
		}

		if r.URL.Path == "/users/opeteer/repos" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode([]map[string]interface{}{
				{
					"name":        "portofolio",
					"description": "Portofolio saya",
					"language":    "Go",
					"html_url":    "https://github.com/opeteer/portofolio",
				},
			})
			return
		}

		http.NotFound(w, r)
	}))
	defer mockServer.Close()

	client := NewClient(2 * time.Second)
	client.BaseURL = mockServer.URL

	report, err := client.FetchLiveTelemetry(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if report.IsFallback {
		t.Errorf("expected live report, got fallback")
	}
	if !report.Probe.Connected {
		t.Errorf("expected probe connected")
	}
	if report.Probe.LatencyMs <= 0 {
		t.Errorf("expected positive latency, got %f", report.Probe.LatencyMs)
	}
	if report.Probe.RateRemain != 55 {
		t.Errorf("expected rate remaining 55, got %d", report.Probe.RateRemain)
	}
	if report.Profile.PublicRepos != 31 {
		t.Errorf("expected 31 repos, got %d", report.Profile.PublicRepos)
	}
	if len(report.RecentRepos) != 1 {
		t.Errorf("expected 1 recent repo, got %d", len(report.RecentRepos))
	}
}

func TestLiveTelemetryFallback(t *testing.T) {
	client := NewClient(100 * time.Millisecond)
	// Invalid unreachable endpoint to test fallback
	client.BaseURL = "http://127.0.0.1:59999"

	report, err := client.FetchLiveTelemetry(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !report.IsFallback {
		t.Errorf("expected fallback report on unreachable endpoint")
	}
	if report.Probe.Connected {
		t.Errorf("expected probe not connected")
	}
	if report.Profile.PublicRepos != 31 {
		t.Errorf("expected 31 repos from fallback dataset, got %d", report.Profile.PublicRepos)
	}
}

package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebRoot(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("expected text/html, got %s", contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Gerardo M Ardianta") {
		t.Errorf("expected body to contain 'Gerardo M Ardianta'")
	}
	if !strings.Contains(body, "emblem.ascii") {
		t.Errorf("expected body to contain 'emblem.ascii'")
	}
}

func TestAPIProfile(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/api/profile", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if res["username"] != "opeteer" {
		t.Errorf("expected username 'opeteer', got '%v'", res["username"])
	}
	if res["email"] != "gmayella245@gmail.com" {
		t.Errorf("expected email 'gmayella245@gmail.com', got '%v'", res["email"])
	}
}

func TestAPIProjects(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/api/projects", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var projects []map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &projects); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if len(projects) != 31 {
		t.Errorf("expected 31 repositories, got %d", len(projects))
	}
}

func TestAPIShowtime(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/api/showtime", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var res map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if len(res["braille"]) == 0 {
		t.Errorf("expected non-empty braille art")
	}
	if len(res["block"]) == 0 {
		t.Errorf("expected non-empty block art")
	}
}

func TestHealthz(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if res["status"] != "healthy" {
		t.Errorf("expected status 'healthy', got '%v'", res["status"])
	}
}

func TestAPILive(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/api/live?limit=3", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	profile, ok := res["profile"].(map[string]interface{})
	if !ok || profile["username"] != "opeteer" {
		t.Errorf("expected profile.username 'opeteer', got '%v'", res["profile"])
	}
	if _, ok := res["probe"].(map[string]interface{}); !ok {
		t.Errorf("expected probe data in response")
	}
}

func TestAPIBench(t *testing.T) {
	handler := NewMux()

	req := httptest.NewRequest("GET", "/api/bench?duration=50ms&target=cpu", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", w.Code)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	sysInfo, ok := res["system_info"].(map[string]interface{})
	if !ok || sysInfo["num_cpu"] == nil {
		t.Errorf("expected system_info in bench response")
	}

	cpu, ok := res["cpu"].(map[string]interface{})
	if !ok || cpu["total_ops"] == nil {
		t.Errorf("expected cpu benchmark data in bench response")
	}
}


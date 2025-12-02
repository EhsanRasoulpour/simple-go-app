package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetingHandler_Default(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	greetingHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 got %d", resp.StatusCode)
	}
	if !strings.Contains(string(body), "Hello") {
		t.Fatalf("expected greeting containing Hello, got: %s", string(body))
	}
}

func TestGreetingHandler_Name(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Tester", nil)
	w := httptest.NewRecorder()
	greetingHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), "Tester") {
		t.Fatalf("expected greeting containing Tester, got: %s", string(body))
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), `"status":"ok"`) && !strings.Contains(string(body), `"status": "ok"`) {
		t.Fatalf("expected JSON with status ok, got %s", string(body))
	}
}

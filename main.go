package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type healthResp struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Started string `json:"started"`
}

var startedAt time.Time

func main() {
	startedAt = time.Now()

	http.HandleFunc("/", greetingHandler)
	http.HandleFunc("/health", healthHandler)

	port := getenv("PORT", "9595")
	addr := fmt.Sprintf(":%s", port)
	log.Printf("starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	greet := getenv("GREETING", "Hello")
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	msg := fmt.Sprintf("%s, %s!", greet, name)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(msg))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startedAt).Round(time.Second).String()
	resp := healthResp{
		Status:  "ok",
		Uptime:  uptime,
		Started: startedAt.Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

type ClockHandler struct{}

type DateTimeResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func (hh ClockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := slog.With("address", r.RemoteAddr)

	// fmt.Println("Remote address:", r.RemoteAddr)
	// fmt.Println("X-Forwarded-For:", r.Header.Get("X-Forwarded-For"))
	// fmt.Println("X-Real-IP", r.Header.Get("X-Real-IP"))

	// Get local time.
	now := time.Now()

	// Get the RFC3339 format string.
	nowStr := now.Format(time.RFC3339)

	// Respond in JSON
	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "application/json") {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		resp := DateTimeResponse{
			DayOfWeek:  now.Weekday().String(),
			DayOfMonth: now.Day(),
			Month:      now.Month().String(),
			Year:       now.Year(),
			Hour:       now.Hour(),
			Minute:     now.Minute(),
			Second:     now.Second(),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.Error(fmt.Sprintf("error writing json: %v", err))
		}
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(fmt.Appendf(nil, "%s\n", nowStr))
	}

	// Log in the JSON format.
	logger.Info("Handled request")
}

func main() {
	// Logger settings
	loggerOptions := &slog.HandlerOptions{Level: slog.LevelDebug}
	logHandler := slog.NewJSONHandler(os.Stderr, loggerOptions)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      ClockHandler{},
	}

	fmt.Println("Serving: http://localhost:8080/")
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

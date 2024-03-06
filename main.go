package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	usersCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "users_served_total",
			Help: "Total number of users served by the server.",
		},
		[]string{"username"},
	)
)

func init() {
	// Register the usersCounter with Prometheus default registry.
	prometheus.MustRegister(usersCounter)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting username from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	username := parts[2]

	// Increment the usersCounter for the given username
	usersCounter.WithLabelValues(username).Inc()

	// Writing response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You are %s\n", username)
}

func main() {
	http.HandleFunc("/user/", userHandler)
	fmt.Println("Server is listening on port 8080...")

	// Expose Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

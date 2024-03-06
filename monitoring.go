package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
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

func RecordUserRequest(username string) {
	usersCounter.WithLabelValues(username).Inc()
}

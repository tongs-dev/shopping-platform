package common

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

// PrometheusBoot sets up a Prometheus metrics endpoint and starts an HTTP server.
func PrometheusBoot(port int) {
	// Validate port number
	if port <= 0 || port > 65535 {
		log.Fatal("Invalid port number: ", port)
	}

	// Register Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Log before starting server
	log.Infof("Prometheus metrics available at: http://0.0.0.0:%d/metrics", port)

	// Start HTTP server in a goroutine
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
		if err != nil {
			log.Fatalf("Failed to start Prometheus metrics server: %v", err)
		}
	}()
}

package common

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// NewTracer initiates a Jaeger Tracer for distributed tracing in a microservices environment using OpenTracing,
// returns tracer instance used for distributed tracing.
func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	if addr == "" {
		addr = "localhost:6831" // Default Jaeger agent address
	}

	// Defines the Jaeger Configuration
	cfg := &jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst, // Sample fractionally
			Param: 0.1,                     // Sample 10% of requests
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}

	// Initializes the Tracer
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Jaeger tracer: %w", err)
	}

	return tracer, closer, nil
}

package monitoring

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// InitTracer inicializa o tracer do OpenTelemetry com um exportador do Jaeger.
func InitTracer() {
	// Configure o exportador do Jaeger
	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost("localhost"),
			jaeger.WithAgentPort("6831"),
		),
	)
	if err != nil {
		log.Fatalf("Failed to create Jaeger exporter: %v", err)
	}

	// Configure o provedor de rastreamento com o exportador do Jaeger
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("minha-api"),
		)),
		trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
}

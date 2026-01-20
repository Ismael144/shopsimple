package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

// Initialize jaeger tracing
func InitTracer(serviceName, jaeger_url string) (func(context.Context) error, error) {
	exp, err := jaeger.New(
		jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint(jaeger_url), 
		),
	)

	if err != nil {
		return nil, err 
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp), 
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL, 
			semconv.ServiceName(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil 
} 
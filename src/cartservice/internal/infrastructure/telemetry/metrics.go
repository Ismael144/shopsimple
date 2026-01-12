package telemetry

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func InitMetrics() (*sdkmetric.MeterProvider, http.Handler, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, nil, err 
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter), 
	)

	otel.SetMeterProvider(provider)

	// exporter expose an HTTP handler via promhttp internally 
	handler := promhttp.Handler()

	return provider, handler, nil 
}
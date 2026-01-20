package interceptors

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/Ismael144/cartservice/internal/infrastructure/requestid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// For logging on every request, lives on interceptor level
func LoggingInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		reqID, _ := requestid.From(ctx)

		// Get span from metadata(headers)
		span := trace.SpanFromContext(ctx)
		traceID := ""

		// If span is not null, then extract the trace id
		if span != nil {
			traceID = span.SpanContext().TraceID().String()
		}

		fields := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.String("request_id", reqID),
			zap.String("trace_id", traceID),
			zap.Duration("duration", time.Since(start)),
		}

		if err != nil {
			log.Error("grpc request failed", append(fields, zap.Error(err))...)
		} else {
			log.Info("grpc request handled", fields...)
		}

		return resp, err
	}
}

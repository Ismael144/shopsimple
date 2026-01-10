package interceptors

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"time"

	"github.com/Ismael144/productservice/internal/infrastructure/requestid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

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

		span := trace.SpanFromContext(ctx)
		traceID := ""

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

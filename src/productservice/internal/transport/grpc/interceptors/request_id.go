package interceptors

import (
	"context"

	"github.com/Ismael144/productservice/internal/infrastructure/requestid"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ctxKey string

const RequestIDKey ctxKey = "request_id"

func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		var reqID string 

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if vals := md.Get("x-request-id"); len(vals) > 0 {
				reqID = vals[0]
			}
		}

		if reqID == "" {
			reqID = uuid.NewString()
		}

		ctx = requestid.With(ctx, reqID)

		// Attach to active span 
		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(
				attribute.String("request.id", reqID),
			)
		}

		// ctx = context.WithValue(ctx, RequestIDKey, uuid.NewString())
		return handler(ctx, req)
	}
}

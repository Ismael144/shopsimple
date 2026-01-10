package interceptors

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ctxKey string

const RequestIDKey ctxKey = "request_id"

func RequestID() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		ctx = context.WithValue(ctx, RequestIDKey, uuid.NewString())
		return handler(ctx, req)
	}
}

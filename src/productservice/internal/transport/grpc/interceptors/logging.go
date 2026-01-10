package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

func Logging() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		resp, err := handler(ctx, req)
		// TODO: We'll add proper structured logging later
		return resp, err
	}
}

package interceptors

import (
	"context"
	"time"

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

		fields := []zap.Field{
			zap.String("method", info.FullMethod), 
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

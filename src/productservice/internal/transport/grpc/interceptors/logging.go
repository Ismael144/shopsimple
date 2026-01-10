package interceptors

import (
	"context"
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

		fields := []zap.Field{
			zap.String("method", info.FullMethod), 
			zap.String("request_id", reqID),
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

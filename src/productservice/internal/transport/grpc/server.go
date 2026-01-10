package grpc

import (
	"net"

	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/productservice/internal/application"
	"github.com/Ismael144/productservice/internal/transport/grpc/handlers"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	// "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpc *grpc.Server
	lis  net.Listener
}

// Start new GRPC Server
func NewServer(
	addr string,
	app *application.ProductService,
	unaryInterceptors ...grpc.UnaryServerInterceptor,
) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	// Adding interceptors
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			unaryInterceptors...
		),
	)

	// Product Service
	productv1.RegisterProductServiceServer(
		server,
		handlers.NewProductHandler(app),
	)

	// Health Service
	healthServer := health.NewServer()
	healthServer.SetServingStatus(
		"", // overall server status
		healthpb.HealthCheckResponse_SERVING,
	)
	healthpb.RegisterHealthServer(server, healthServer)

	// Reflection
	reflection.Register(server)

	// Serve GRPC Server
	return &Server{
		grpc: server,
		lis:  lis,
	}, nil
}

// Start server
func (s *Server) Start() error {
	return s.grpc.Serve(s.lis)
}

// Stop server
func (s *Server) Stop() {
	s.grpc.GracefulStop()
}

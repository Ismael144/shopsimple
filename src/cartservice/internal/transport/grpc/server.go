package grpc

import (
	"net"
	"time"

	cartv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/cart/v1"
	"github.com/Ismael144/cartservice/internal/application"
	"github.com/Ismael144/cartservice/internal/transport/grpc/clients"
	"github.com/Ismael144/cartservice/internal/transport/grpc/handlers"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpc *grpc.Server 
	lis net.Listener
}

func NewServer(
	addr string, 
	cart_service *application.CartService, 
	product_client *clients.ProductClient, 
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
			unaryInterceptors...,
		), 
		grpc.ConnectionTimeout(5*time.Second),
	)

	// Cart service 
	cartv1.RegisterCartServiceServer(
		server,
		handlers.NewCartHandler(cart_service, product_client),
	)

	// Health service 
	healthServer := health.NewServer()
	healthServer.SetServingStatus(
		"",
		healthpb.HealthCheckResponse_SERVING,
	)
	healthpb.RegisterHealthServer(server, healthServer)

	// Reflection 
	reflection.Register(server)

	// Serve GRPC Server
	return &Server{
		grpc: server, 
		lis: lis,
	}, nil 
}

func (s *Server) Start() error {
	return s.grpc.Serve(s.lis)
}

func (s *Server) Stop() {
	s.grpc.GracefulStop()
}
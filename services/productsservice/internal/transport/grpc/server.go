package grpc

import (
	"net"

	productv1 "github.com/Ismael144/productsservice/gen/go/proto/product/v1"
	"github.com/Ismael144/productsservice/internal/application"
	"github.com/Ismael144/productsservice/internal/transport/grpc/handlers"
	"google.golang.org/grpc"
)

type Server struct {
	grpc *grpc.Server
	lis  net.Listener
}

// Start new GRPC Server
func NewServer(
	addr string,
	app *application.ProductsService,
	unaryInterceptors ...grpc.UnaryServerInterceptor,
) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	// Adding interceptors
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
	)

	// Register Product Service
	productv1.RegisterProductServiceServer(
		server,
		handlers.NewProductHandler(app),
	)

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

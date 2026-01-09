package grpc

import (
	"net"

	productv1 "github.com/Ismael144/productsservice/gen/go/proto/product/v1"
	"github.com/Ismael144/productsservice/internal/application"
	"github.com/Ismael144/productsservice/internal/transport/grpc/handlers"
	"google.golang.org/grpc"
)

// Start new GRPC Server
func NewServer(
	addr string,
	app *application.ProductsService,
	unaryInterceptors ...grpc.UnaryServerInterceptor,
) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
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
	return server.Serve(lis)
}

package clients

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	productv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/cartservice/internal/transport/grpc/interceptors"
)

type ProductClient struct {
	Client productv1.ProductServiceClient
}

func NewProductClient(addr string) (*ProductClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), 
		// Add client request deadline of 2 seconds 
		grpc.WithUnaryInterceptor(interceptors.WithDefaultTimeout(2 * time.Second)), 
	)
	
	if err != nil {
		return nil, err
	}
	// defer conn.Close()
	
	client := productv1.NewProductServiceClient(conn)
	
	return &ProductClient{
		Client: client,
	}, nil 
}

package clients

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	productv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/product/v1"

	"github.com/Ismael144/cartservice/internal/transport/grpc/interceptors"
)

// Product Service Server Client
type ProductServiceClient struct {
	client productv1.ProductServiceClient
}

// Initialize ProductService Server Client
func NewProductServiceServerClient(addr string) (*ProductServiceClient, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Add client request deadline of 2 seconds
		grpc.WithUnaryInterceptor(interceptors.WithDefaultTimeout(2*time.Second)),
	)

	if err != nil {
		return nil, err
	}

	client := productv1.NewProductServiceClient(conn)

	return &ProductServiceClient{
		client: client,
	}, nil
}

// Getter for product service client
func (p *ProductServiceClient) GetClient() productv1.ProductServiceClient {
	return p.client
}

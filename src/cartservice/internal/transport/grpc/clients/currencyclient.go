package clients

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	currencyv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/currency/v1"
	
	"github.com/Ismael144/cartservice/internal/transport/grpc/interceptors"
)

type CurrencyServiceClient struct {
	client currencyv1.CurrencyServiceClient
}

// Initialize CurrencyService Server Client
func NewCurrencyServiceServerClient(addr string) (*CurrencyServiceClient, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Add client request deadline of 2 seconds
		grpc.WithUnaryInterceptor(interceptors.WithDefaultTimeout(2*time.Second)),
	)

	if err != nil {
		return nil, err
	}

	client := currencyv1.NewCurrencyServiceClient(conn)

	return &CurrencyServiceClient{
		client: client,
	}, nil
}

// Getter for currency service client
func (c *CurrencyServiceClient) GetClient() currencyv1.CurrencyServiceClient {
	return c.client
}

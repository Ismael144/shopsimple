package handlers

import (
	"context"

	common "github.com"
	cartv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/cart/v1"
	productv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/cartservice/internal/application"
	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
	"github.com/Ismael144/cartservice/internal/transport/grpc/clients"

	// "github.com/Ismael144/cartservice/internal/transport/grpc/mapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartHandler struct {
	cartv1.UnimplementedCartServiceServer
	service        *application.CartService
	product_client *clients.ProductClient
}

func NewCartHandler(service *application.CartService, product_client *clients.ProductClient) *CartHandler {
	return &CartHandler{service: service, product_client: product_client}
}

func (c *CartHandler) AddItem(ctx context.Context, req *cartv1.AddToCartRequest) (*cartv1.AddToCartResponse, error) {
	product, err := c.product_client.Client.FindById(ctx, &productv1.FindByIdRequest{
		Id: req.ProductId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	cartItem := entities.CartItem{
		ProductID:   valueobjects.ParseProductID(product.Id),
		ProductName: product.Name,
		Quantity:    req.Quantity,
		UnitPrice:   common.MoneyFromCents(int64(product.UnitPrice)),
	}

	productId, err := c.service.AddToCart(ctx, valueobjects.UserID(req.UserId), &cartItem)
	if err != nil {
		return nil, err
	}

	return &cartv1.AddToCartResponse{
		ProductId: productId.String(),
	}, nil
}

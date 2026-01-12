package handlers

import (
	"context"

	cartv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/cart/v1"
	productv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/cartservice/internal/application"
	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
	"github.com/Ismael144/cartservice/internal/transport/grpc/clients"
	"github.com/Ismael144/cartservice/internal/transport/grpc/mapper"

	// "github.com/Ismael144/cartservice/internal/transport/grpc/mapper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartHandler struct {
	cartv1.UnimplementedCartServiceServer
	service       *application.CartService
	productclient *clients.ProductClient
}

func NewCartHandler(service *application.CartService, product_client *clients.ProductClient) *CartHandler {
	return &CartHandler{service: service, productclient: product_client}
}

func (c *CartHandler) AddItem(ctx context.Context, req *cartv1.ModifyCartRequest) (*cartv1.ModifyCartResponse, error) {
	cartItem, err := c.service.IsItemInCart(ctx, valueobjects.UserID(req.UserId), valueobjects.ProductID(req.ProductId))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// If item is not in cart, we then make grpc request to product service
	if cartItem == nil {
		product, err := c.productclient.Client.FindById(ctx, &productv1.FindByIdRequest{
			Id: req.ProductId,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cartItem := entities.CartItem{
			ProductID:   valueobjects.ParseProductID(product.Id),
			ProductName: product.Name,
			Quantity:    req.Quantity,
			UnitPrice:   mapper.FromProtoMoney(product.UnitPrice),
		}

		_, err = c.service.AddToCart(ctx, valueobjects.ParseUserID(req.UserId), &cartItem)
		if err != nil {
			return nil, err
		}
	} else {
		// If cartItem exits in db, We construct a cart item without having to make a grpc request to productservice
		cartItem := entities.CartItem{
			ProductID: valueobjects.ParseProductID(req.ProductId),
			Quantity:  req.Quantity,
			ProductName: cartItem.ProductName,
			UnitPrice: cartItem.UnitPrice,
		}

		_, err := c.service.AddToCart(ctx, valueobjects.ParseUserID(req.UserId), &cartItem)
		if err != nil {
			return nil, err
		}
	}

	return &cartv1.ModifyCartResponse{
		ProductId: req.ProductId,
	}, nil
}

func (c *CartHandler) DeductFromCart(ctx context.Context, req *cartv1.ModifyCartRequest) (*cartv1.ModifyCartResponse, error) {
	productID, err := c.service.DeductFromCart(ctx, valueobjects.UserID(req.UserId), valueobjects.ProductID(req.ProductId), req.Quantity)

	if err != nil {
		return nil, err
	}

	return &cartv1.ModifyCartResponse{
		ProductId: productID.String(),
	}, nil
}

func (c *CartHandler) GetCart(ctx context.Context, req *cartv1.GetCartRequest) (*cartv1.GetCartResponse, error) {
	cart, err := c.service.GetCart(ctx, valueobjects.UserID(req.UserId))
	if err != nil {
		return nil, err
	}

	return &cartv1.GetCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

func (c *CartHandler) RemoveFromCart(ctx context.Context, req *cartv1.RemoveFromCartRequest) (*cartv1.ModifyCartResponse, error) {
	productID, err := c.service.RemoveFromCart(ctx, valueobjects.ParseUserID(req.UserId), valueobjects.ParseProductID(req.ProductId))
	if err != nil {
		return nil, err 
	}

	return &cartv1.ModifyCartResponse{
		ProductId: productID.String(),
	}, nil 
} 

func (c *CartHandler) Clear(ctx context.Context, req *cartv1.GetCartRequest) (*cartv1.ModifyCartResponse, error) {
	err := c.service.Clear(ctx, valueobjects.ParseUserID(req.UserId))
	if err != nil {
		return nil, err
	}
	return &cartv1.ModifyCartResponse{
		ProductId: req.UserId,
	}, nil 
}

func (c *CartHandler) AssignToAuthUser(ctx context.Context, req *cartv1.AssignToAuthUserRequest) (*cartv1.AssignToAuthUserResponse, error) {
	err := c.service.AssignToAuthUser(ctx, valueobjects.ParseUserID(req.GuestUserId), valueobjects.ParseUserID(req.AuthUserId))
	if err != nil {
		return nil, err
	}
	return &cartv1.AssignToAuthUserResponse{
		AuthUserId: req.AuthUserId,
	}, nil 
}
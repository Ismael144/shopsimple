package handlers

import (
	"context"

	cartv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/cart/v1"
	commonv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/common/v1"
	currencyv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/currency/v1"
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

// Implements gRPC Cart Service Endpoints
type CartHandler struct {
	cartv1.UnimplementedCartServiceServer
	service        *application.CartService
	productclient  *clients.ProductServiceClient
	currencyclient *clients.CurrencyServiceClient
}

func NewCartHandler(service *application.CartService, product_client *clients.ProductServiceClient, currency_client *clients.CurrencyServiceClient) *CartHandler {
	return &CartHandler{service: service, productclient: product_client, currencyclient: currency_client}
}

// GRPC Method for AddItem on CartService
func (c *CartHandler) AddItem(ctx context.Context, req *cartv1.ModifyCartRequest) (*cartv1.ModifyCartResponse, error) {
	cart := &entities.Cart{}

	// Make findbyid grpc call to productservice inorder to get product details
	product, err := c.productclient.GetClient().FindById(ctx, &productv1.FindByIdRequest{
		Id: req.ProductId,
	})
	// Err handling
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Build cartitem depending on details of the product
	cartItem := entities.CartItem{
		ProductID:   valueobjects.ParseProductID(product.Id),
		ProductName: product.Name,
		Quantity:    req.Quantity,
		UnitPrice:   mapper.FromProtoMoney(product.UnitPrice),
	}

	// Add cart item to cart, with productStock,
	// A call is made to productservice on every
	// AddToCart command, for stock validation.
	cart, err = c.service.AddToCart(ctx, valueobjects.ParseUserID(req.UserId), &cartItem, uint32(product.Stock))
	if err != nil {
		return nil, err
	}

	return &cartv1.ModifyCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

// gRPC Method DeductFromCart for CartService
func (c *CartHandler) DeductFromCart(ctx context.Context, req *cartv1.ModifyCartRequest) (*cartv1.ModifyCartResponse, error) {
	// Deduct quantity from cart item
	cart, err := c.service.DeductFromCart(ctx, valueobjects.UserID(req.UserId), valueobjects.ProductID(req.ProductId), req.Quantity)

	if err != nil {
		return nil, err
	}

	return &cartv1.ModifyCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

func (c *CartHandler) GetCart(ctx context.Context, req *cartv1.GetCartRequest) (*cartv1.GetCartResponse, error) {
	// Get current state of cart with precomputed cart price total
	cart, err := c.service.GetCart(ctx, valueobjects.UserID(req.UserId))
	if err != nil {
		return nil, err
	}

	return &cartv1.GetCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

// GRPC Method RemoveFromCart for CartService
func (c *CartHandler) RemoveFromCart(ctx context.Context, req *cartv1.RemoveFromCartRequest) (*cartv1.ModifyCartResponse, error) {
	// Remove item from cart
	cart, err := c.service.RemoveFromCart(ctx, valueobjects.ParseUserID(req.UserId), valueobjects.ParseProductID(req.ProductId))
	if err != nil {
		return nil, err
	}

	return &cartv1.ModifyCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

// GRPC Method Clear for CartService
func (c *CartHandler) Clear(ctx context.Context, req *cartv1.GetCartRequest) (*cartv1.ModifyCartResponse, error) {
	// Empty cart
	err := c.service.EmptyCart(ctx, valueobjects.ParseUserID(req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// Get current state of cart
	cart, err := c.service.GetCart(ctx, valueobjects.UserID(req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &cartv1.ModifyCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

// GRPC Method AssignToAuthUser for CartService
func (c *CartHandler) AssignToAuthUser(ctx context.Context, req *cartv1.AssignToAuthUserRequest) (*cartv1.AssignToAuthUserResponse, error) {
	err := c.service.AssignToAuthUser(ctx, valueobjects.ParseUserID(req.GuestUserId), valueobjects.ParseUserID(req.AuthUserId))
	if err != nil {
		return nil, err
	}
	return &cartv1.AssignToAuthUserResponse{
		AuthUserId: req.AuthUserId,
	}, nil
}

// Converts cart items currency to selected currency by user
func (c *CartHandler) ConvertToCurrency(ctx context.Context, req *cartv1.ConvertToCurrencyRequest) (*cartv1.ModifyCartResponse, error) {
	// List of cart items money values
	var cartItemsMoneyValues []*commonv1.Money

	// Get current state of the cart with precomputed price totals
	cart, err := c.service.GetCart(ctx, valueobjects.ParseUserID(req.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Collect the money values of all cart items
	for _, cartItem := range cart.Items {
		cartItemsMoneyValues = append(cartItemsMoneyValues, mapper.ToProtoMoney(cartItem.UnitPrice))
	}
	// Form BatchCurrencyConversionRequest
	batchConversionReq := currencyv1.BatchCurrencyConversionRequest{
		FromMoney: cartItemsMoneyValues,
		ToCode:    req.Currency,
	}
	// Make grpc call to currency service, returning a list
	// of converted money values
	convertedMoneyValues, err := c.currencyclient.GetClient().BatchConvert(ctx, &batchConversionReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// Handle error from currency
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// Update cart with new converted money
	for i := range len(convertedMoneyValues.Results) {
		cart.Items[i].UnitPrice = mapper.FromProtoMoney(convertedMoneyValues.Results[i])
	}
	// After updating cart, save the currently modified cart into db
	err = c.service.Save(ctx, valueobjects.ParseUserID(req.UserId), cart)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &cartv1.ModifyCartResponse{
		Cart: mapper.ToProtoCart(cart),
	}, nil
}

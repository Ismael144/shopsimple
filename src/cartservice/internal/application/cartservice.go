package application

import (
	"context"

	"github.com/Ismael144/cartservice/internal/application/ports"
	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

type CartService struct {
	cartrepo ports.CartRepository
}

func NewCartService(cartrepo ports.CartRepository) *CartService {
	return &CartService{cartrepo: cartrepo}
}

// Add item to cart, with stock validation
// Returns an ErrStockLimitReached if product
// stock is exceeded, productStock will not be
// stored in db because it will lead to stock
// inconsistencies
func (c *CartService) AddToCart(ctx context.Context, userID valueobjects.UserID, item *entities.CartItem, productStock uint32) (*entities.Cart, error) {
	return c.cartrepo.AddItem(ctx, userID, item, productStock)
}

// Deduct quantity from product in cart
func (c *CartService) DeductFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID, Quantity uint32) (*entities.Cart, error) {
	return c.cartrepo.DeductFromCart(ctx, userID, productID, Quantity)
}

// This is not a true service, will be used to check whether product is in cart already
// If so, then it does not call the product service since it has product details from first call
func (c *CartService) IsItemInCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.CartItem, error) {
	return c.cartrepo.IsItemInCart(ctx, userID, productID)
}

// Remove item from cart, by specifying user id and product id
func (c *CartService) RemoveFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.Cart, error) {
	return c.cartrepo.RemoveFromCart(ctx, userID, productID)
}

// Assign guest user cart to authenticated user
func (c *CartService) AssignToAuthUser(ctx context.Context, guestUserID valueobjects.UserID, authUserID valueobjects.UserID) error {
	return c.cartrepo.AssignToAuthUser(ctx, guestUserID, authUserID)
}

// Empty cart
func (c *CartService) EmptyCart(ctx context.Context, userID valueobjects.UserID) error {
	return c.cartrepo.EmptyCart(ctx, userID)
}

// Return current state of the cart with
// pre computed cart total
func (c *CartService) GetCart(ctx context.Context, userID valueobjects.UserID) (*entities.Cart, error) {
	return c.cartrepo.GetCart(ctx, userID)
}

// Save current state of cart
func (c *CartService) Save(ctx context.Context, userID valueobjects.UserID, cart *entities.Cart) error {
	return c.cartrepo.Save(ctx, userID, cart)
}

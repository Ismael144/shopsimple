package ports

import (
	"context"

	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

// CartRepository interface, containing all cartservice db operations
type CartRepository interface {
	// Helper function for checking existence of item in cart
	IsItemInCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.CartItem, error)
	GetCart(ctx context.Context, userID valueobjects.UserID) (*entities.Cart, error)
	AddItem(ctx context.Context, userID valueobjects.UserID, cartItem *entities.CartItem, productStock uint32) (*entities.Cart, error)
	DeductFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID, Quantity uint32) (*entities.Cart, error)
	RemoveFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.Cart, error)
	AssignToAuthUser(ctx context.Context, guestUserID valueobjects.UserID, authUserID valueobjects.UserID) error
	EmptyCart(ctx context.Context, userID valueobjects.UserID) error
	Save(ctx context.Context, userID valueobjects.UserID, cart *entities.Cart) error
}

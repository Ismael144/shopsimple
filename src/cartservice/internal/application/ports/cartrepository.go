package ports

import (
	"context"

	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

// CartRepository interface, containing all cartservice db operations
type CartRepository interface {
	GetCart(ctx context.Context, userID valueobjects.UserID) (*entities.Cart, error)
	IsItemInCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.CartItem, error)
	AddItem(ctx context.Context, userID valueobjects.UserID, cartItem *entities.CartItem) (*valueobjects.ProductID, error)
	DeductFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID, Quantity uint32) (*valueobjects.ProductID, error)
	RemoveFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*valueobjects.ProductID, error)
	AssignToAuthUser(ctx context.Context, guestUserID valueobjects.UserID, authUserID valueobjects.UserID) error
	Clear(ctx context.Context, userID valueobjects.UserID) error
}

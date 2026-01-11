package ports

import (
	"context"

	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

type CartRepository interface {
	AddItem(ctx context.Context, userID valueobjects.UserID, cartItem *entities.CartItem) (valueobjects.ProductID, error)
	GetCart(ctx context.Context, userID valueobjects.UserID) (entities.Cart, error)
}

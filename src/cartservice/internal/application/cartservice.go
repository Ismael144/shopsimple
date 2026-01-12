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

func (c *CartService) AddToCart(ctx context.Context, userID valueobjects.UserID, item *entities.CartItem) (*valueobjects.ProductID, error) {
	return c.cartrepo.AddItem(ctx, userID, item)
} 
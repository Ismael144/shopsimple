package application

import "github.com/Ismael144/cartservice/internal/application/ports"

type CartService struct {
	cartrepo ports.CartRepository
}

func NewCartService(cartrepo ports.CartRepository) *CartService {
	return &CartService{cartrepo: cartrepo}
}
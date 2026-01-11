package entities

import "github.com/Ismael144/cartservice/internal/domain/valueobjects"

type CartItem struct {
	ID          valueobjects.CartID
	ProductID   valueobjects.ProductID
	Quantity    uint32
	ProductName string
	UnitPrice   valueobjects.Money
}

func NewCartItem(ProductID string, Quantity uint32, ProductName string, UnitPrice valueobjects.Money) CartItem {
	return CartItem{
		ID:          valueobjects.NewCartID(),
		ProductID:   valueobjects.ParseProductID(ProductID),
		ProductName: ProductName,
		Quantity:    Quantity,
		UnitPrice:   UnitPrice,
	}
}

package entities

import (
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

// Cart item domain 
type CartItem struct {
	ProductID   valueobjects.ProductID `json:"product_id"`
	Quantity    uint32                 `json:"quantity"`
	ProductName string                 `json:"product_name"`
	UnitPrice   valueobjects.Money     `json:"unitprice"`
}

// Initialize cart item
func NewCartItem(ProductID string, Quantity uint32, ProductName string, UnitPrice valueobjects.Money) CartItem {
	return CartItem{
		ProductID:   valueobjects.ParseProductID(ProductID),
		ProductName: ProductName,
		Quantity:    Quantity,
		UnitPrice:   UnitPrice,
	}
}

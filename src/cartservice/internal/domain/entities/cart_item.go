package entities

import (
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
	"github.com"
)

type CartItem struct {
	ProductID   valueobjects.ProductID `json:"product_id"`
	Quantity    uint32                 `json:"quantity"`
	ProductName string                 `json:"product_name"`
	UnitPrice   common.Money     `json:"unitprice"`
}

func NewCartItem(ProductID string, Quantity uint32, ProductName string, UnitPrice common.Money) CartItem {
	return CartItem{
		ProductID:   valueobjects.ParseProductID(ProductID),
		ProductName: ProductName,
		Quantity:    Quantity,
		UnitPrice:   UnitPrice,
	}
}

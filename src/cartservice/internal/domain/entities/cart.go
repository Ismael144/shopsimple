package entities

import (
	"slices"

	"github.com/Ismael144/cartservice/internal/domain"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

type Cart struct {
	UserID         valueobjects.UserID
	Items          []*CartItem
	cartTotal valueobjects.Money
}

func NewCart(UserID valueobjects.UserID) *Cart {
	return &Cart{
		UserID:         UserID,
		Items:          []*CartItem{},
		cartTotal: valueobjects.MoneyFromCents(0),
	}
}

func (cart *Cart) AddToCart(CartItem *CartItem) error {
	if CartItem.Quantity <= 0 {
		return domain.ErrInvalidQuantity
	}

	for i, existing := range cart.Items {
		if existing.ProductID == CartItem.ProductID {
			cart.Items[i].Quantity += CartItem.Quantity
			return nil
		}
	}

	cart.Items = append(cart.Items, CartItem)

	return nil
}

func (cart *Cart) Total() valueobjects.Money {
	total := valueobjects.Money{Cents: 0}
	for _, item := range cart.Items {
		total = total.Add(item.UnitPrice.Mul(int64(item.Quantity)))
	}

	return total
}

func (cart *Cart) RemoveItem(productID valueobjects.ProductID) {
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items = slices.Delete(cart.Items, i, i+1)
			return
		}
	}
}

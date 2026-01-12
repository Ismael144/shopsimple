package entities

import (
	"encoding/json"
	"slices"

	"github.com"
	"github.com/Ismael144/cartservice/internal/domain"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

type Cart struct {
	UserID valueobjects.UserID `json:"user_id"`
	Items  []*CartItem         `json:"items"`
}

func NewCart(UserID valueobjects.UserID) *Cart {
	return &Cart{
		UserID: UserID,
		Items:  []*CartItem{},
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

func (cart *Cart) DeductFromCart(productID valueobjects.ProductID, Quantity uint32) {
	for i, existing := range cart.Items {
		if existing.ProductID == productID {
			if existing.Quantity <= Quantity {
				cart.RemoveItem(existing.ProductID)
			} else {
				cart.Items[i].Quantity -= Quantity
			}
		}
	}
}

func (cart *Cart) Total() common.Money {
	total := common.Money{Cents: 0}
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

func UnmarshalCart(cartJson string) (*Cart, error) {
	var cart Cart
	err := json.Unmarshal([]byte(cartJson), &cart)
	return &cart, err
}

func (cart *Cart) Marshal() (string, error) {
	cartJson, err := json.Marshal(cart)
	if err != nil {
		return "", err
	}
	return string(cartJson), nil
}

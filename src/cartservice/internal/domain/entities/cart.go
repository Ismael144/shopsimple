package entities

import (
	"encoding/json"
	"slices"

	"github.com/Ismael144/cartservice/internal/domain"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

// Cart domain
type Cart struct {
	UserID     valueobjects.UserID `json:"user_id"`
	Items      []*CartItem         `json:"items"`
	PriceTotal valueobjects.Money
}

func NewCart(UserID valueobjects.UserID) *Cart {
	return &Cart{
		UserID:     UserID,
		Items:      []*CartItem{},
		PriceTotal: valueobjects.MoneyFromCents(0),
	}
}

func (cart *Cart) GetById(productId valueobjects.ProductID) *CartItem {
	for _, item := range cart.Items {
		if item.ProductID == productId {
			return item
		}
	}
	return nil
}

func (cart *Cart) AddToCart(item *CartItem) error {
	// Some qty validation, its uint32, meaning this validation check is completely useless
	// I'ma jus leave it there :)
	if item.Quantity <= 0 {
		return domain.ErrInvalidQuantity
	}
	cartItem := cart.GetById(item.ProductID)
	if cartItem != nil {
		cartItem.Quantity += item.Quantity
	} else {
		cart.Items = append(cart.Items, item)
	}

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
			return
		}
	}
}

// Compute cart price total
func (cart *Cart) Total() valueobjects.Money {
	total := valueobjects.Money{Cents: 0}
	for _, item := range cart.Items {
		total = total.Add(item.UnitPrice.Mul(int64(item.Quantity)))
	}
	cart.PriceTotal = total

	return total
}

// Remove item by id from cart
func (cart *Cart) RemoveItem(productID valueobjects.ProductID) {
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items = slices.Delete(cart.Items, i, i+1)
			return
		}
	}
}

// Clear cart
func (cart *Cart) Clear() {
	cart.Items = []*CartItem{}
}

// Computes total and saves the total in Total field in cart
func (cart *Cart) GetCart() *Cart {
	cart.Total()
	return cart
}

// Cart helper functions

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

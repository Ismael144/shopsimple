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
		// Currency code will be determined by first item
		// added to cart
		PriceTotal: valueobjects.NewMoney("", 0, 0),
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

func (cart *Cart) AddToCart(item *CartItem, productStock uint32) (*Cart, error) {
	// Some qty validation, its uint32, meaning this validation check is completely useless
	// I'ma jus leave it there :)
	if item.Quantity == 0 {
		return nil, domain.ErrInvalidQuantity
	}
	cartItem := cart.GetById(item.ProductID)
	// Stock validation
	if item.Quantity > productStock {
		return nil, domain.ErrStockLimitReached
	}
	if cartItem != nil {
		cartItem.Quantity += item.Quantity
	} else {
		cart.Items = append(cart.Items, item)
	}

	return cart.GetCart(), nil
}

func (cart *Cart) DeductFromCart(productID valueobjects.ProductID, Quantity uint32) *Cart {
	for i, existing := range cart.Items {
		if existing.ProductID == productID {
			if existing.Quantity <= Quantity {
				cart.RemoveItem(existing.ProductID)
			} else {
				cart.Items[i].Quantity -= Quantity
			}
			break
		}
	}
	return cart.GetCart()
}

// Compute cart price total
func (cart *Cart) Total() valueobjects.Money {
	total := valueobjects.NewMoney("", 0, 0)
	for _, item := range cart.Items {
		// We set currency code of total variable basing 
		// on currency of first item in cart
		total.CurrencyCode = item.UnitPrice.GetCurrencyCode()
		total = valueobjects.Must(
			valueobjects.Sum(
				total, valueobjects.Multiply(item.UnitPrice, item.Quantity),
			),
		)
	}
	cart.PriceTotal = total

	return total
}

// Remove item by id from cart
func (cart *Cart) RemoveItem(productID valueobjects.ProductID) *Cart {
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items = slices.Delete(cart.Items, i, i+1)
			break
		}
	}
	return cart.GetCart()
}

// Clear cart
func (cart *Cart) Empty() *Cart {
	cart.Items = []*CartItem{}
	return cart.GetCart()
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

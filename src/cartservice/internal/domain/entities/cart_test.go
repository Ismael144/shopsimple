package entities

import (
	"testing"

	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

func InitCartWithItems() *Cart {
	cart := NewCart(valueobjects.UserID("some-uuid"))
	
	cartItem1 := NewCartItem("1", 4, "some item1", valueobjects.Dollars(100))
	cartItem2 := NewCartItem("2", 4, "some item2", valueobjects.Dollars(100))
	cartItem3 := NewCartItem("3", 4, "some item3", valueobjects.Dollars(100))

	cart.AddToCart(&cartItem1)
	cart.AddToCart(&cartItem2)
	cart.AddToCart(&cartItem3)

	return cart
}

func TestAddToCart(t *testing.T) {
	cart := InitCartWithItems()

	if len(cart.Items) != 3 {
		t.Errorf("Expected cart length: 3 found: %d", len(cart.Items))
	}
}

func TestCartTotal(t *testing.T) {
	cart := InitCartWithItems()
	if cart.Total().String() != "$1200.00" {
		t.Errorf("Expected total: 1200.00, found: %s", cart.Total().String())
	}
}

func TestRemoveCart(t *testing.T) {
	cart := InitCartWithItems()
	cart.RemoveItem("1")
	for _, item := range cart.Items {
		if item.ProductID.String() == "1" {
			t.Errorf("Could not delete item from cart")
		}
	}

	if len(cart.Items) != 2 { 
		t.Errorf("Expected cart length: 2, found %d", len(cart.Items)) 
	}
}

func TestSimple(t *testing.T) {
	t.Errorf("Simple test")
}

package entities

import (
	"testing"

	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

func InitCartWithItems() *Cart {
	cart := NewCart(valueobjects.UserID("some-uuid"), "USD")

	cartItem1 := NewCartItem("1", 4, "some item1", valueobjects.NewMoney("USD", 100, 29))
	cartItem2 := NewCartItem("2", 4, "some item2", valueobjects.NewMoney("USD", 100, 29))
	cartItem3 := NewCartItem("3", 4, "some item3", valueobjects.NewMoney("USD", 100, 29))

	cart.AddToCart(&cartItem1, 10)
	cart.AddToCart(&cartItem2, 10)
	cart.AddToCart(&cartItem3, 10)

	return cart
}

func TestProuctItemStockLimitReached(t *testing.T) {
	cart := InitCartWithItems()
	cartItem := NewCartItem("3", 5, "some item3", valueobjects.NewMoney("USD", 100, 29))
	_, err := cart.AddToCart(&cartItem, 4)

	if err == nil {
		t.Errorf("Expected ErrStockLimitReach Error, but found: %v", err)
	}
}

func TestAddToCart(t *testing.T) {
	cart := InitCartWithItems()

	if len(cart.Items) != 3 {
		t.Errorf("Expected cart length: 3 found: %d", len(cart.Items))
	}
}

func TestGetById(t *testing.T) {
	cart := InitCartWithItems()

	cartitem := cart.GetById(valueobjects.ParseProductID("1"))
	if cartitem == nil {
		t.Errorf("Expected cart item")
	}
}

// Testing cart where id of all products is the same
func TestAddToCartWithSameIds(t *testing.T) {
	cart := NewCart(valueobjects.UserID("some-uuid"))

	cartItem1 := NewCartItem("1", 4, "some item1", valueobjects.NewMoney("USD", 100, 29))
	cartItem2 := NewCartItem("1", 4, "some item2", valueobjects.NewMoney("USD", 100, 29))
	cartItem3 := NewCartItem("1", 4, "some item3", valueobjects.NewMoney("USD", 100, 29))

	cart.AddToCart(&cartItem1, 2)
	cart.AddToCart(&cartItem2, 2)
	cart.AddToCart(&cartItem3, 2)

	if len(cart.Items) != 1 {
		t.Errorf("Expected cart length: 1 found: %d", len(cart.Items))
	}
}

func TestCartTotal(t *testing.T) {
	cart := InitCartWithItems()
	if cart.Total().String() != "$1200.00" {
		t.Errorf("Expected total: 1200.00, found: %s", cart.Total().String())
	}
}

func TestDeductFromCart(t *testing.T) {
	cart := InitCartWithItems()
	cart.DeductFromCart(valueobjects.ParseProductID("1"), 2)
	if cart.Items[0].Quantity != 2 {
		t.Errorf("Expected 2, got: %d", cart.Items[0].Quantity)
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

func TestCartMarshal(t *testing.T) {
	cart := InitCartWithItems()
	_, err := cart.Marshal()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCartUnmarshal(t *testing.T) {
	cart := InitCartWithItems()
	// Marshal cart
	_, err := cart.Marshal()
	if err != nil {
		t.Errorf("An error occured: %v", err)
	}
}

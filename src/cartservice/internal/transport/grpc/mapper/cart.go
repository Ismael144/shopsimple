package mapper

import (
	"fmt"

	cartv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/cart/v1"
	commonv1 "github.com/Ismael144/cartservice/gen/go/shopsimple/common/v1"
	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
)

func ToProtoMoney(m valueobjects.Money) *commonv1.Money {
	return &commonv1.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}
}

// Convert proto money value type to money value domain
func FromProtoMoney(m *commonv1.Money) valueobjects.Money {
	return valueobjects.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}
}

// Convert cart item domain to cart item proto type
func ToProtoCartItem(cartItemDomain *entities.CartItem) cartv1.CartItem {
	return cartv1.CartItem{
		ProductId:   cartItemDomain.ProductID.String(),
		ProductName: cartItemDomain.ProductName,
		Quantity:    cartItemDomain.Quantity,
		UnitPrice:   ToProtoMoney(cartItemDomain.UnitPrice),
	}
}

// Convert cart domain type to cart proto type
func ToProtoCart(cartDomain *entities.Cart) *cartv1.Cart {
	fmt.Println(cartDomain)
	cartItems := make([]*cartv1.CartItem, 0, len(cartDomain.Items))
	for _, item := range cartDomain.Items {
		cartItem := ToProtoCartItem(item)
		cartItems = append(cartItems, &cartItem)
	}

	return &cartv1.Cart{
		UserId:    cartDomain.UserID.String(),
		CartItems: cartItems,
	}
}

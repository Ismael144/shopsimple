package valueobjects

import "github.com/google/uuid"

type CartID string

func NewCartID() CartID {
	return CartID(uuid.NewString())
}

func ParseCartID(id string) CartID {
	return CartID(id)
}

func (u CartID) String() string {
	return u.String()
}

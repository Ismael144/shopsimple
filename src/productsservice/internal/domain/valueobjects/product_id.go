package valueobjects

import "github.com/google/uuid"

type ProductID string

func NewProductID() ProductID {
	return ProductID(uuid.NewString())
}

func ParseProductID(id string) ProductID {
	return ProductID(id)
}

func (id ProductID) String() string {
	return string(id)
}

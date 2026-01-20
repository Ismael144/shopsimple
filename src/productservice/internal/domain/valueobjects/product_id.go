package valueobjects

import "github.com/google/uuid"

type ProductID string

// Init Product Id 
func NewProductID() ProductID {
	return ProductID(uuid.NewString())
}

// Parse String Based UUID 
func ParseProductID(id string) ProductID {
	return ProductID(id)
}

// String representative of id
func (id ProductID) String() string {
	return string(id)
}

package valueobjects

type ProductID string

func ParseProductID(id string) ProductID {
	return ProductID(id)
}

func (u ProductID) String() string {
	return string(u)
}

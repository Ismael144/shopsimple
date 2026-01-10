package valueobjects

import "github.com/google/uuid"

type CategoryID string

func NewCategoryID() CategoryID {
	return CategoryID(uuid.NewString())
}

func ParseCategoryID(id string) CategoryID {
	return CategoryID(id)
}

func (id CategoryID) String() string {
	return string(id)
}

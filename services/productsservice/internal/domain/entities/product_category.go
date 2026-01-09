package domain

import (
	"time"

	"github.com/Ismael144/productsservice/internal/domain/valueobjects"
)

type ProductCategory struct {
	ID        valueobjects.CategoryID
	Name      string
	CreatedAt time.Time
}

func NewProductCategory(Name string, Now time.Time) ProductCategory {
	return ProductCategory{
		ID:        valueobjects.NewCategoryID(),
		Name:      Name,
		CreatedAt: Now,
	}
}

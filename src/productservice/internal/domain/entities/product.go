package domain

import (
	"time"

	"github.com/Ismael144/productservice/internal/domain/valueobjects"
)

type Product struct {
	ID          valueobjects.ProductID
	ProductName string
	Description string
	UnitPrice   float64
	ImageUrl    string
	Stock       int64
	CategoryID  valueobjects.CategoryID
	CreatedAt   time.Time
}

func NewProduct(product_name, description string, unit_price float64, image_url string, stock int64, category_id string, now time.Time) Product {
	return Product{
		ID:          valueobjects.NewProductID(),
		ProductName: product_name,
		Description: description,
		UnitPrice:   unit_price,
		ImageUrl:    image_url,
		Stock:       stock,
		CategoryID:  valueobjects.ParseCategoryID(category_id),
		CreatedAt:   now,
	}
}

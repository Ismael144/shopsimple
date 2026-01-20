package domain

import (
	"time"

	"github.com/Ismael144/productservice/internal/domain/valueobjects"
)

type Product struct {
	ID          valueobjects.ProductID
	Name        string
	Description string
	UnitPrice   valueobjects.Money
	ImageUrl    string
	Stock       int64
	Rating      uint32
	// A list of categories
	Categories []string
	CreatedAt  time.Time
}

// Initialize new product, takes in money from database
func NewProduct(name, description string, unit_price valueobjects.Money, image_url string, stock int64, categories []string, rating uint32, now time.Time) Product {
	return Product{
		ID:          valueobjects.NewProductID(),
		Name:        name,
		Description: description,
		UnitPrice:   unit_price,
		ImageUrl:    image_url,
		Stock:       stock,
		Rating:      rating,
		Categories:  categories,
		CreatedAt:   now,
	}
}

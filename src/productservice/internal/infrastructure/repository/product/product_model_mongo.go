package product

import (
	"time"
)

// This is mongodb's database model
type ProductModelMongo struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	PriceMinor  int64  `bson:"price_minor"`
	Currency    string `bson:"currency"`
	ImageUrl    string `bson:"image_url"`
	Stock       int64  `bson:"stock"`

	// Refering to product rating, limit 5
	rating int32 `bson:"rating"`

	// A list of categories a given
	// product is under
	Categories []string  `bson:"categories"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

package product

import (
	"time"
	product_category "github.com/Ismael144/productservice/internal/infrastructure/repository/product_category"
)

type ProductModel struct {
	ID          string  `gorm:"primaryKey;type:uuid;"`
	Name        string  `gorm:"uniqueIndex;not null"`
	Description string  `gorm:"null"`
	UnitPrice   float64 `gorm:"not null"`
	ImageUrl    string
	Stock       int64                `gorm:"not null"`
	CategoryId  string               `gorm:"column:category_id;foreignKey:CategoryID;references:ID"`
	Category    product_category.ProductCategoryModel `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ProductModel) TableName() string {
	return "products"
}

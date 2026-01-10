package product_category

import "time"

type ProductCategoryModel struct {
	ID   string `gorm:"primaryKey;type:uuid"`
	Name string
	CreatedAt time.Time
}

func (ProductCategoryModel) TableName() string {
	return "product_categories"
}
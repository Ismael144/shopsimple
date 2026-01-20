package product_category

import "time"

type ProductCategoryModelPg struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	Name      string
	CreatedAt time.Time
}

func (ProductCategoryModelPg) TableName() string {
	return "product_categories"
}

package repository

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	ID          string               `gorm:"primaryKey;type:uuid;"`
	ProductName string               `gorm:"uniqueIndex;not null"`
	Description string               `gorm:"null"`
	UnitPrice   float64              `gorm:"not null"`
	Stock       int64                `gorm:"not null"`
	CategoryId  uint                 `gorm:"column:category_id;foreignKey:CategoryID;references:ID"`
	Category    ProductCategoryModel `gorm:"foreignKey:CategoryID"`
}

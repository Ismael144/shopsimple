package repository

import "gorm.io/gorm"

type ProductCategoryModel struct {
	gorm.Model
	ID   string `gorm:"primaryKey;type:uuid"`
	Name string
}

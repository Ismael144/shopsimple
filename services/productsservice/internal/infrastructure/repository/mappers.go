package repository

import (
	"time"

	domain "github.com/Ismael144/productsservice/internal/domain/entities"
	"github.com/Ismael144/productsservice/internal/domain/valueobjects"
)

func productModelToDomain(m *ProductModel) *domain.Product {
	return &domain.Product{
		ID:          valueobjects.ProductID(m.ID),
		ProductName: m.ProductName,
		Description: m.Description,
		UnitPrice:   m.UnitPrice,
		ImageUrl:    m.ImageUrl,
		Stock:       m.Stock,
		CategoryID:  valueobjects.CategoryID(m.Category.ID),
		CreatedAt:   m.CreatedAt,
	}
}

func productDomainToModel(u *domain.Product) ProductModel {
	return ProductModel{
		ID:          u.ID.String(),
		ProductName: u.ProductName,
		Description: u.Description,
		UnitPrice:   u.UnitPrice,
		ImageUrl:    u.ImageUrl,
		Stock:       u.Stock,
		CategoryId:  u.CategoryID.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func categoryModelToDomain(m *ProductCategoryModel) domain.ProductCategory {
	return domain.ProductCategory{
		ID:        valueobjects.CategoryID(m.ID),
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func categoryDomainToModel(m *ProductCategoryModel) ProductCategoryModel {
	return ProductCategoryModel{
		ID:   m.ID,
		Name: m.Name,
	}
}

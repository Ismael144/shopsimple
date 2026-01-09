package repository

import (
	"strconv"

	domain "github.com/Ismael144/productsservice/internal/domain/entities"
	"github.com/Ismael144/productsservice/internal/domain/valueobjects"
)

func productModelToDomain(m *ProductModel) *domain.Product {
	return &domain.Product{
		ID:          valueobjects.ProductID(m.ID),
		ProductName: m.ProductName,
		Description: m.Description,
		UnitPrice:   m.UnitPrice,
		Stock:       m.Stock,
		CategoryID:  valueobjects.CategoryID(m.Category.ID),
		CreatedAt:   m.CreatedAt,
	}
}

func productDomainToModel(u *domain.Product) ProductModel {
	category_id, _ := strconv.ParseUint(u.CategoryID.String(), 10, 64)

	return ProductModel{
		ID:          u.ID.String(),
		ProductName: u.ProductName,
		Description: u.Description,
		UnitPrice:   u.UnitPrice,
		Stock:       u.Stock,
		CategoryId:  uint(category_id),
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

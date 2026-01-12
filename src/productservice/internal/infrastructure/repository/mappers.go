package repository

import (
	"time"

	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"github.com/Ismael144/productservice/internal/infrastructure/repository/product"
	"github.com/Ismael144/productservice/internal/infrastructure/repository/product_category"
)

func ProductModelToDomain(m *product.ProductModel) *domain.Product {
	return &domain.Product{
		ID:          valueobjects.ProductID(m.ID),
		Name:        m.Name,
		Description: m.Description,
		UnitPrice:   valueobjects.MoneyFromCents(int64(m.UnitPrice)),
		ImageUrl:    m.ImageUrl,
		Stock:       m.Stock,
		CategoryID:  valueobjects.CategoryID(m.Category.ID),
		CreatedAt:   m.CreatedAt,
	}
}

func ProductDomainToModel(u *domain.Product) product.ProductModel {
	return product.ProductModel{
		ID:          u.ID.String(),
		Name:        u.Name,
		Description: u.Description,
		UnitPrice:   float64(u.UnitPrice.Cents),
		ImageUrl:    u.ImageUrl,
		Stock:       u.Stock,
		CategoryId:  u.CategoryID.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func CategoryModelToDomain(m *product_category.ProductCategoryModel) *domain.ProductCategory {
	return &domain.ProductCategory{
		ID:        valueobjects.CategoryID(m.ID),
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}

func CategoryDomainToModel(m *domain.ProductCategory) product_category.ProductCategoryModel {
	return product_category.ProductCategoryModel{
		ID:	m.ID.String(),
		Name: m.Name,
	}
}

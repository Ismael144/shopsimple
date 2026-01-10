package repository

import (
	"context"

	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{db: db}
}

func (r *ProductCategoryRepository) Create(ctx context.Context, p *domain.ProductCategory) error {
	model := categoryDomainToModel(p)
	return r.db.WithContext(ctx).
		Create(model).
		Error
}

func (r *ProductCategoryRepository) List(ctx context.Context) ([]*domain.ProductCategory, uint32, error) {
	var (
		models     []*ProductCategoryModel
		totalCount int64
	)

	query := r.db.WithContext(ctx).
		Model(&ProductCategoryModel{})

	if err := query.Session(&gorm.Session{}).Count(&totalCount).Error; err != nil {
		return nil, 0, nil
	}

	if err := query.Find(&models).Error; err != nil {
		return nil, 0, nil
	}

	// Convert ProductCategoryModel to ProductCategory domian
	domainCategories := make([]*domain.ProductCategory, 0, len(models))
	for _, productCategory := range models {
		domainCategories = append(domainCategories, categoryModelToDomain(productCategory))
	}

	return domainCategories, uint32(totalCount), nil
}

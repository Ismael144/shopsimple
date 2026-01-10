package repository

import (
	"context"

	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (r *ProductsRepository) Create(ctx context.Context, product *domain.Product) error {
	model := productDomainToModel(product)
	return r.db.WithContext(ctx).
		Create(model).
		Error
}

func (r *ProductsRepository) UpdateStock(ctx context.Context, product_id *valueobjects.ProductID, stock int64) error {
	return r.db.WithContext(ctx).
		Where("id = ?", product_id.String()).
		Update("stock = ?", stock).
		Error
}

func (r *ProductsRepository) FindById(ctx context.Context, product_id *valueobjects.ProductID) (*domain.Product, error) {
	m := ProductModel{}
	err := r.db.WithContext(ctx).
		Where("id = ?", product_id.String()).
		First(&m).
		Error

	if err != nil {
		return nil, err
	}

	// Convert ProductModel to ProductDomain
	p := productModelToDomain(&m)
	return p, nil
}

func (r *ProductsRepository) List(ctx context.Context, page, pageSize uint32) ([]*domain.Product, uint32, error) {
	var (
		models []ProductModel
		total  int64
	)

	query := r.db.WithContext(ctx).
		Model(&ProductModel{})

	if err := query.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Safety check
	if page < 1 {
		page = 1
	}
	offset := int((page - 1) * pageSize)

	err := query.Limit(int(pageSize)).
		Offset(offset).
		Find(&models).
		Error

	if err != nil {
		return nil, 0, err
	}

	products := make([]*domain.Product, 0, len(models))
	for _, product := range models {
		products = append(products, productModelToDomain(&product))
	}

	return products, uint32(total), nil
}

func (r *ProductsRepository) FilterByProductFiltersObject(ctx context.Context, pf *ports.ProductFilters) ([]*domain.Product, uint32, error) {
	var models []*ProductModel
	var totalCount int64

	query := r.db.WithContext(ctx).
		Model(&ProductModel{})

	if pf.SearchString != "" {
		query = query.Where("NAME LIKE ?", "%"+pf.SearchString+"%")
	}

	if len(pf.Categories) != 0 {
		query = query.Where("category_id IN ?", pf.Categories)
	}

	// Nested price filter
	if pf.Prices != nil {
		if pf.Prices.Min > 0 {
			query = query.Where("unit_price >= ?", pf.Prices.Min)
		}

		if pf.Prices.Max > 0 {
			query = query.Where("unit_price <= ?", pf.Prices.Max)
		}
	}

	// Get total count
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if pf.PageSize > 0 {
		offset := (pf.Page - 1) * pf.PageSize
		query = query.Limit(int(pf.PageSize)).Offset(int(offset))
	}

	// Execute final query
	if err := query.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	products := make([]*domain.Product, 0, len(models))
	for _, product_model := range models {
		products = append(products, productModelToDomain(product_model))
	}

	return products, uint32(totalCount), nil
}

package repository

import (
	"context"

	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"github.com/Ismael144/productservice/internal/infrastructure/repository/product"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (r *ProductsRepository) Create(ctx context.Context, product *domain.Product) error {
	model := ProductDomainToModel(product)
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
	m := product.ProductModel{}
	err := r.db.WithContext(ctx).
		Where("id = ?", product_id.String()).
		First(&m).
		Error

	if err != nil {
		return nil, err
	}

	// Convert ProductModel to ProductDomain
	p := ProductModelToDomain(&m)
	return p, nil
}

func (r *ProductsRepository) List(ctx context.Context, page, pageSize uint32) ([]*domain.Product, uint32, error) {
	var (
		models []product.ProductModel
		total  int64
	)

	query := r.db.WithContext(ctx).
		Model(&product.ProductModel{})

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
		products = append(products, ProductModelToDomain(&product))
	}

	return products, uint32(total), nil
}

func (r *ProductsRepository) Filter(ctx context.Context, pf *ports.ProductFilters) ([]*domain.Product, uint32, error) {
	var models []*product.ProductModel
	var totalCount int64

	query := r.db.WithContext(ctx).
		Model(&product.ProductModel{})

	if pf.SearchString != "" {
		query = query.Where("name LIKE ?", "%"+pf.SearchString+"%")
	}

	if len(pf.Categories) != 0 {
		query = query.Where("category_id IN ?", pf.Categories)
	}

	// Nested price filter
	if pf.PriceRanges != nil {
		if pf.PriceRanges.Min > 0 {
			query = query.Where("unit_price >= ?", pf.PriceRanges.Min)
		}

		if pf.PriceRanges.Max > 0 {
			query = query.Where("unit_price <= ?", pf.PriceRanges.Max)
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
		products = append(products, ProductModelToDomain(product_model))
	}

	return products, uint32(totalCount), nil
}

func (r *ProductsRepository) BatchFindById(ctx context.Context, product_ids []*valueobjects.ProductID) ([]*domain.Product, uint32, error) {
	var (
		models     []product.ProductModel
		totalCount int64
	)
	// Convert product ids to string
	product_ids_string := make([]string, 0, len(product_ids))
	for _, product_id := range product_ids {
		product_ids_string = append(product_ids_string, product_id.String())
	}

	query := r.db.WithContext(ctx).
		Model(&product.ProductModel{})
		
	// Find multiple items by id in db
	query = query.Find(&models, "id IN ?", product_ids_string)

	// Get row count
	if err := query.Session(&gorm.Session{}).Count(&totalCount).Error; err != nil {
		return nil, 0, nil
	}

	products := make([]*domain.Product, 0, len(models))
	for _, product_model := range models {
		products = append(products, ProductModelToDomain(&product_model))
	}

	return products, uint32(totalCount), nil
}

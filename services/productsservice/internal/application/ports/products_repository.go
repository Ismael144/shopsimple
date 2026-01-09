package repository

import (
	"context"

	"github.com/Ismael144/productsservice/internal/domain/entities"
	"github.com/Ismael144/productsservice/internal/domain/valueobjects"
)

type PricesFilter struct {
	min int64
	max int64
}

type ProductFilters struct {
	SearchString string
	Categories   []*valueobjects.CategoryID
	Prices       *PricesFilter
}

type ProductsRespository interface {
	List(ctx context.Context, page, pageSize uint32) ([]*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	IncreaseStock(ctx context.Context, stock int64) error
	DecreaseStock(ctx context.Context, stock int64) error
	FindById(ctx context.Context, product_id *valueobjects.ProductID) (*domain.Product, error)
	FilterByObject(ctx context.Context, product_filters *ProductFilters) ([]*domain.Product, error)
}

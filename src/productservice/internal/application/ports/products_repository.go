package ports

import (
	"context"

	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
)

type PriceRanges struct {
	Min float64
	Max float64
}

type ProductFilters struct {
	Page         uint32
	PageSize     uint32
	SearchString string
	Categories   []valueobjects.CategoryID
	PriceRanges  *PriceRanges
}

type ProductsRespository interface {
	Create(ctx context.Context, product *domain.Product) error
	List(ctx context.Context, page, pageSize uint32) ([]*domain.Product, uint32, error)
	UpdateStock(ctx context.Context, product *valueobjects.ProductID, stock int64) error
	FindById(ctx context.Context, product_id *valueobjects.ProductID) (*domain.Product, error)
	FilterByProductFiltersObject(ctx context.Context, product_filters *ProductFilters) ([]*domain.Product, uint32, error)
}

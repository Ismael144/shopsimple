package ports

import (
	"context"

	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
)

type PriceRanges struct {
	Min valueobjects.Money
	Max valueobjects.Money
}

// Used to apply filters to products
type ProductFilters struct {
	Page         uint32
	PageSize     uint32
	SearchString string
	Categories   []string
	PriceRanges  *PriceRanges
}

// Defines how the structure of the product repository
// Where any database could be used and still return the same results
// or be predictable
type ProductsRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	List(ctx context.Context, page, pageSize uint32) ([]*domain.Product, *domain.Pagination, error)
	UpdateStock(ctx context.Context, product *valueobjects.ProductID, stock int64) error
	FindById(ctx context.Context, product_id *valueobjects.ProductID) (*domain.Product, error)
	Filter(ctx context.Context, product_filters *ProductFilters) ([]*domain.Product, *domain.Pagination, error)
	BatchFindById(ctx context.Context, product_ids []*valueobjects.ProductID) ([]*domain.Product, int64, error)
	ListCategories(ctx context.Context) ([]string, error)
}

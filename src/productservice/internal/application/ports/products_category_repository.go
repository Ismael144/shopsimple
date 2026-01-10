package ports

import (
	"context"

	domain "github.com/Ismael144/productservice/internal/domain/entities"
)

type ProductCategoryRepository interface {
	Create(ctx context.Context, category *domain.ProductCategory) error
	List(ctx context.Context) ([]*domain.ProductCategory, uint32, error)
}

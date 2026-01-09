package ports

import (
	"context"

	domain "github.com/Ismael144/productsservice/internal/domain/entities"
)

type ProductCategory interface {
	List(ctx context.Context) ([]*domain.Product, error)
	Create(ctx context.Context, category *domain.ProductCategory) error
}

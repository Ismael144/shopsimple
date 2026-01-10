package application

import (
	"context"

	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
)

type ProductCategoryService struct {
	repo ports.ProductCategoryRepository
}

func NewProductCategoryService(repo ports.ProductCategoryRepository) *ProductCategoryService {
	return &ProductCategoryService{repo: repo}
}

func (p *ProductCategoryService) Create(ctx context.Context, new_category *domain.ProductCategory) error {
	return p.repo.Create(ctx, new_category)
}

func (p *ProductCategoryService) List(ctx context.Context) ([]*domain.ProductCategory, uint32, error) {
	return p.repo.List(ctx)
}
package application

import (
	"context"

	"github.com/Ismael144/productsservice/internal/application/ports"
	domain "github.com/Ismael144/productsservice/internal/domain/entities"
)

type ProductsService struct {
	repo ports.ProductsRespository
}

func NewProductsService(repo ports.ProductsRespository) *ProductsService {
	return &ProductsService{repo: repo}
}

func (p *ProductsService) List(ctx context.Context, page, pageSize uint32) ([]*domain.Product, uint32, error) {
	return p.repo.List(ctx, page, pageSize)
}

func (p *ProductsService) Create(ctx context.Context, new_product *domain.Product) error {
	return p.repo.Create(ctx, new_product)
}

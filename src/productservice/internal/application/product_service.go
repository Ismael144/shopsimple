package application

import (
	"context"

	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
)

type ProductService struct {
	repo ports.ProductsRespository
}

func Newproductservice(repo ports.ProductsRespository) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) List(
	ctx context.Context,
	page,
	pageSize uint32,
) ([]*domain.Product, uint32, error) {
	return p.repo.List(ctx, page, pageSize)
}

func (p *ProductService) Create(
	ctx context.Context,
	new_product *domain.Product,
) error {
	return p.repo.Create(ctx, new_product)
}

func (p *ProductService) Filter(
	ctx context.Context,
	product_filters *ports.ProductFilters,
) ([]*domain.Product, uint32, error) {
	return p.repo.FilterByProductFiltersObject(ctx, product_filters)
}

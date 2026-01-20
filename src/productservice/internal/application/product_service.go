package application

import (
	"context"

	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
)

type ProductService struct {
	repo ports.ProductsRepository
}

func NewProductService(repo ports.ProductsRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) List(
	ctx context.Context,
	page,
	pageSize uint32,
) ([]*domain.Product, *domain.Pagination, error) {
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
) ([]*domain.Product, *domain.Pagination, error) {
	return p.repo.Filter(ctx, product_filters)
}

func (p *ProductService) FindById(
	ctx context.Context,
	product_id string,
) (*domain.Product, error) {
	return p.repo.FindById(ctx, (*valueobjects.ProductID)(&product_id))
}

func (p *ProductService) BatchFindById(
	ctx context.Context,
	product_ids []*valueobjects.ProductID,
) ([]*domain.Product, int64, error) {
	return p.repo.BatchFindById(ctx, product_ids)
}

func (p *ProductService) ListCategories(
	ctx context.Context, 
) ([]string, error) {
	return p.repo.ListCategories(ctx)
}
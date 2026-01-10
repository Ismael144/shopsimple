package handlers

import (
	"context"

	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/productservice/internal/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductCategoryHandler struct {
	productv1.UnimplementedProductServiceServer
	categories *application.ProductCategoryService
}

func NewProductCategoryHandler(categories *application.ProductCategoryService) *ProductCategoryHandler {
	return &ProductCategoryHandler{categories: categories}
}

func (c *ProductCategoryHandler) Create(ctx context.Context, req *productv1.CreateCategoryRequest) (*productv1.CreateCategoryResponse, error) {
	categoryDomain := productv1.CreateCategoryRequest{
		Name: req.Name,
	}

	if err := c.categories.Create(ctx, categoryDomain); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.CreateCategoryResponse{}, nil
}

func (c *ProductCategoryHandler) List(ctx context.Context, req *productv1.ListCategoryRequest) (*productv1.ListCategoryResponse, error) {
	categories, totalCount, err := c.categories.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert domain categories to grpc categories
	grpcCategories := make([]*productv1.ProductCategory, 0, len(categories))
	for _, category := range categories {
		grpcCategories = append(grpcCategories, &productv1.ProductCategory{
			Name: category.Name,
		})
	}

	return &productv1.ListCategoryResponse{
		Categories: grpcCategories,
		Total:      totalCount,
	}, nil
}

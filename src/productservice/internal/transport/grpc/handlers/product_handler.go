package handlers

import (
	"context"
	"fmt"
	"time"

	_ "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/productservice/internal/application"
	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"github.com/Ismael144/productservice/internal/transport/grpc/mapper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductHandler struct {
	productv1.UnimplementedProductServiceServer
	products   *application.ProductService
	categories *application.ProductCategoryService
}

func NewProductHandler(products *application.ProductService, categories *application.ProductCategoryService) *ProductHandler {
	return &ProductHandler{products: products, categories: categories}
}

func (h *ProductHandler) List(ctx context.Context, req *productv1.ListRequest) (*productv1.ListResponse, error) {
	products, totalCount, err := h.products.List(ctx, req.Page, req.PageSize)

	// Check error existence
	if err != nil {
		return &productv1.ListResponse{
			Products: []*productv1.Product{},
			Total:    0,
		}, status.Error(codes.Internal, err.Error())
	}

	// Convert domain products to proto products
	protoProducts := mapper.ToProtoProducts(products)

	return &productv1.ListResponse{
		Products: protoProducts,
		Total:    totalCount,
	}, nil
}

func (h *ProductHandler) Create(ctx context.Context, req *productv1.CreateRequest) (*productv1.CreateResponse, error) {
	newProduct := &domain.Product{
		ID:          valueobjects.NewProductID(),
		Name:        req.Name,
		Description: req.Description,
		UnitPrice:   float64(req.UnitPrice),
		Stock:       int64(req.Stock),
		CategoryID:  valueobjects.CategoryID(req.CategoryId),
		CreatedAt:   time.Now(),
	}

	if err := h.products.Create(ctx, newProduct); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.CreateResponse{}, nil
}

func (h *ProductHandler) Filter(ctx context.Context, req *productv1.FilterRequest) (*productv1.ListResponse, error) {
	fmt.Println(req.Page, req.PageSize, req.SearchString, req.Categories, req.PriceRanges)
	// Convert category ids string to CategoryID type
	categories := make([]valueobjects.CategoryID, 0, len(req.Categories))
	for _, categoryID := range req.Categories {
		categories = append(categories, valueobjects.ParseCategoryID(categoryID))
	}

	// Convert repo.ProductFilters to grpc.ProductFilters
	product_filters := ports.ProductFilters{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Categories: categories,
		PriceRanges: &ports.PriceRanges{
			Min: float64(req.PriceRanges.Min),
			Max: float64(req.PriceRanges.Max),
		},
	}

	products, totalCount, err := h.products.Filter(ctx, &product_filters)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoProducts := mapper.ToProtoProducts(products)

	return &productv1.ListResponse{
		Products: protoProducts,
		Total:    totalCount,
	}, nil
}

// Product Categories Services 

func (c *ProductHandler) CreateCategory(ctx context.Context, req *productv1.CreateCategoryRequest) (*productv1.CreateCategoryResponse, error) {
	categoryDomain := domain.NewProductCategory(req.Name, time.Now())

	if err := c.categories.Create(ctx, &categoryDomain); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.CreateCategoryResponse{}, nil
}

func (c *ProductHandler) ListCategories(ctx context.Context, req *productv1.ListCategoriesRequest) (*productv1.ListCategoriesResponse, error) {
	categories, totalCount, err := c.categories.List(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert domain categories to grpc categories
	grpcCategories := make([]*productv1.ProductCategory, 0, len(categories))
	for _, category := range categories {
		grpcCategories = append(grpcCategories, &productv1.ProductCategory{
			Id: category.ID.String(),
			Name: category.Name,
			CreatedAt: timestamppb.New(category.CreatedAt),
		})
	}

	return &productv1.ListCategoriesResponse{
		Categories: grpcCategories,
		Total:      totalCount,
	}, nil
}

package handlers

import (
	"context"
	"fmt"
	"time"

	_ "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	commonv1 "github.com/Ismael144/productservice/gen/go/shopsimple/common/v1"
	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	"github.com/Ismael144/productservice/internal/application"
	"github.com/Ismael144/productservice/internal/application/ports"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"github.com/Ismael144/productservice/internal/transport/grpc/mapper"
)

type ProductHandler struct {
	productv1.UnimplementedProductServiceServer
	products *application.ProductService
}

// Initialize product handler
func NewProductHandler(products *application.ProductService) *ProductHandler {
	return &ProductHandler{products: products}
}

// List available products with pagination
func (h *ProductHandler) List(ctx context.Context, req *productv1.ListRequest) (*productv1.ListResponse, error) {
	products, pagination, err := h.products.List(ctx, req.Page, req.PageSize)

	// Check error existence
	if err != nil {
		return &productv1.ListResponse{
			Products:   []*productv1.Product{},
			Pagination: nil,
		}, status.Error(codes.Internal, err.Error())
	}

	// Convert domain products to proto products
	protoProducts := mapper.ToProtoProducts(products)
	fmt.Println(protoProducts)

	return &productv1.ListResponse{
		Products:   protoProducts,
		Pagination: mapper.ToProtoPagination(pagination),
	}, nil
}

// Create a new product
func (h *ProductHandler) Create(ctx context.Context, req *productv1.CreateRequest) (*productv1.CreateResponse, error) {
	newProduct := &domain.Product{
		ID:          valueobjects.NewProductID(),
		Name:        req.Name,
		Description: req.Description,
		UnitPrice:   mapper.FromProtoMoney(req.UnitPrice),
		Stock:       int64(req.Stock),
		Categories:  req.Categories,
		CreatedAt:   time.Now(),
	}

	if err := h.products.Create(ctx, newProduct); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.CreateResponse{}, nil
}

// Filter products by properties
func (h *ProductHandler) Filter(ctx context.Context, req *productv1.FilterRequest) (*productv1.ListResponse, error) {
	// Convert category ids string to CategoryID type
	categories := req.Categories

	// We initialize price ranges
	// This is done coz if req.PriceRanges is not provided
	// when calling the service and you try to access req.PriceRange
	// leading to a "invalid memory address or nil pointer dereference" error
	var (
		PriceRangeMin = &commonv1.Money{
			CurrencyCode: "USD", 
			Units: 0, 
			Nanos: 0,
		}
		PriceRangeMax = &commonv1.Money{
			CurrencyCode: "USD", 
			Units: 0, 
			Nanos: 0,
		}
	)

	// We check whether req.PriceRanges is not null 
	// to prevent null pointer dereference, if not null,
	// we assign the min and max of price ranges
	if req.PriceRanges != nil {
		PriceRangeMin = req.PriceRanges.Min
		PriceRangeMax = req.PriceRanges.Max
	}

	// Convert repo.ProductFilters to grpc.ProductFilters
	product_filters := ports.ProductFilters{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Categories: categories,
		PriceRanges: &ports.PriceRanges{
			Min: mapper.FromProtoMoney(PriceRangeMin),
			Max: mapper.FromProtoMoney(PriceRangeMax),
		},
	}

	products, pagination, err := h.products.Filter(ctx, &product_filters)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoProducts := mapper.ToProtoProducts(products)

	return &productv1.ListResponse{
		Products:   protoProducts,
		Pagination: mapper.ToProtoPagination(pagination),
	}, nil
}

// Take in a list of multiple product ids to return their details 
func (h *ProductHandler) BatchFindById(ctx context.Context, req *productv1.BatchFindByIdRequest) (*productv1.BatchFindByIdResponse, error) {
	// Convert BatchFindByIdRequest into ProductIDs
	product_ids := make([]*valueobjects.ProductID, 0, len(req.Ids))
	for _, id := range req.Ids {
		productID := valueobjects.ParseProductID(id)
		product_ids = append(product_ids, &productID)
	}

	products, totalCount, err := h.products.BatchFindById(ctx, product_ids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Convert domain Products to Proto Products
	protoProducts := mapper.ToProtoProducts(products)

	return &productv1.BatchFindByIdResponse{
		Products: protoProducts,
		Total:    totalCount,
	}, nil
}

// Find Product details by id 
func (h *ProductHandler) FindById(ctx context.Context, req *productv1.FindByIdRequest) (*productv1.Product, error) {
	product, err := h.products.FindById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return mapper.ToProtoProduct(product), nil
}

func (h *ProductHandler) ListCategories(ctx context.Context, req *emptypb.Empty) (*productv1.ListCategoriesResponse, error) {
	categories, err := h.products.ListCategories(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.ListCategoriesResponse{
		Categories: categories,
	}, nil 
}
package mapper

import (
	commonv1 "github.com/Ismael144/productservice/gen/go/shopsimple/common/v1"
	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoMoney(m valueobjects.Money) *commonv1.Money {
	return &commonv1.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}
}

// Convert proto money value type to money value domain
func FromProtoMoney(m *commonv1.Money) valueobjects.Money {
	return valueobjects.Money{
		CurrencyCode: m.CurrencyCode,
		Units:        m.Units,
		Nanos:        m.Nanos,
	}
}

// Convert domain product to grpc product
func ToProtoProduct(product *domain.Product) *productv1.Product {
	return &productv1.Product{
		Id:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		UnitPrice:   ToProtoMoney(product.UnitPrice),
		ImageUrl:    product.ImageUrl,
		Stock:       uint64(product.Stock),
		Categories:  product.Categories,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}
}

// Convert a list of domain products to grpc products
func ToProtoProducts(products []*domain.Product) []*productv1.Product {
	grpcProducts := make([]*productv1.Product, 0, len(products))
	for _, product := range products {
		grpcProducts = append(grpcProducts, ToProtoProduct(product))
	}
	return grpcProducts
}

// Convert pagination object into proto pagination object
func ToProtoPagination(pagination *domain.Pagination) *commonv1.Pagination {
	return &commonv1.Pagination{
		CurrentPage: pagination.CurrentPage,
		TotalItems: pagination.TotalItems,
		TotalPages: pagination.TotalPages,
	}
}
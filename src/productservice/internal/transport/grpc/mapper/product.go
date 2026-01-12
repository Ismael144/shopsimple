package mapper

import (
	commonv1 "github.com/Ismael144/productservice/gen/go/shopsimple/common/v1"
	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoMoney(money valueobjects.Money) *commonv1.Money {
	return &commonv1.Money{
		Cents: money.Cents,
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
		CategoryId:  product.CategoryID.String(),
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

package mapper

import (
	productv1 "github.com/Ismael144/productservice/gen/go/shopsimple/product/v1"
	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoProduct(product *domain.Product) *productv1.Product {
	return &productv1.Product{
		Id:          product.ID.String(),
		ProductName: product.ProductName,
		Description: product.Description,
		UnitPrice:   product.UnitPrice,
		ImageUrl:    product.ImageUrl,
		Stock:       uint64(product.Stock),
		CategoryId:  product.CategoryID.String(),
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}
}

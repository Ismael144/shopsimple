package mapper

import (
	userv1 "github.com/Ismael144/productsservice/gen/go/proto/product/v1"
	domain "github.com/Ismael144/productsservice/internal/domain/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoProduct(product *domain.Product) *userv1.Product {
	return &userv1.Product{
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

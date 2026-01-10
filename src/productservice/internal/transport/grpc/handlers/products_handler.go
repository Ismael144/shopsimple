package handlers

import (
	"context"

	_ "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	productv1 "github.com/Ismael144/productservice/gen/go/proto/product/v1"
	"github.com/Ismael144/productservice/internal/application"
	"github.com/Ismael144/productservice/internal/transport/grpc/mapper"
)

type ProductHandler struct {
	productv1.UnimplementedProductServiceServer
	products *application.ProductService
}

func NewProductHandler(products *application.ProductService) *ProductHandler {
	return &ProductHandler{products: products}
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
	protoProducts := []*productv1.Product{}
	for _, product := range products {
		protoProducts = append(protoProducts, mapper.ToProtoProduct(product))
	}

	return &productv1.ListResponse{
		Products: protoProducts,
		Total:    totalCount,
	}, nil
}

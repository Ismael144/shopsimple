package handlers

import (
	"context"

	"google.golang.org/grpc"

	userv1 "github.com/Ismael144/productsservice/gen/go/proto/product/v1"
	"github.com/Ismael144/productsservice/internal/application"
)

type ProductHandler struct {
	userv1.UnimplementedProductServiceServer
	products *application.ProductsService
}

func NewProductHandler(products *application.ProductsService) *ProductHandler {
	return &ProductHandler{products: products}
}

func (h *ProductHandler) List(ctx context.Context, req *userv1.ListRequest) (*userv1.ListResponse, error) {
	// Unimplemeneted
}
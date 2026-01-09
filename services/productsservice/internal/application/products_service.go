package application
import (
	_ "context"

	"github.com/Ismael144/productsservice/internal/application/ports"
)

type ProductsService struct {
	repo *ports.ProductsRespository
}

func NewProductsService(repo *ports.ProductsRespository) *ProductsService {
	return &ProductsService{repo: repo}
}
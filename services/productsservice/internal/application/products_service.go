package application
import (
	_ "context"

	"github.com/Ismael144/productsservice/internal/application/ports"
)

type ProductsService struct {
	repo *repository.ProductsRespository
}

func NewProductsService(repo *repository.ProductsRespository) *ProductsService {
	return &ProductsService{repo: repo}
}
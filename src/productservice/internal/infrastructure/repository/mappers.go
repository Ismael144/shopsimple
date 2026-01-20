package repository

import (
	"time"

	domain "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"github.com/Ismael144/productservice/internal/infrastructure/repository/product"
)

func moneyToMinor(m valueobjects.Money) int64 {
	return m.Units*1_000_000_000 + int64(m.Nanos)
}

func minorToMoney(minor int64, currency string) valueobjects.Money {
	return valueobjects.Money{
		CurrencyCode: currency,
		Units:        minor / 1_000_000_000,
		Nanos:        int32(minor % 1_000_000_000),
	}
}

// Convert product model to product domain
func ProductModelToDomain(m *product.ProductModelMongo) *domain.Product {
	return &domain.Product{
		ID:          valueobjects.ProductID(m.ID),
		Name:        m.Name,
		Description: m.Description,
		UnitPrice:   minorToMoney(m.PriceMinor, m.Currency),
		ImageUrl:    m.ImageUrl,
		Stock:       m.Stock,
		Categories:  m.Categories,
		CreatedAt:   m.CreatedAt,
	}
}

// Convert product domain to model
func ProductDomainToModel(d *domain.Product) product.ProductModelMongo {
	return product.ProductModelMongo{
		ID:          d.ID.String(),
		Name:        d.Name,
		Description: d.Description,
		PriceMinor:  moneyToMinor(valueobjects.NewMoney(d.UnitPrice.GetCurrencyCode(), d.UnitPrice.GetUnits(), d.UnitPrice.GetNanos())),
		Currency:    d.UnitPrice.GetCurrencyCode(),
		ImageUrl:    d.ImageUrl,
		Stock:       d.Stock,
		Categories:  d.Categories,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

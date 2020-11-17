package postgres

import (
	"context"

	"github.com/andrefebrianto/rest-api/src/models"
)

// ProductCommand represent command function for product domain
type ProductCommand interface {
	CreateProduct(context context.Context, product *models.Product) ([]models.Product, error)
	UpdateProduct(context context.Context, product *models.Product) (*models.Product, error)
	DeleteProduct(context context.Context, id string) error
}

// ProductQueries represent query function for product domain
type ProductQueries interface {
	GetProducts(context context.Context, limit, page int) ([]models.Product, error)
	GetProductByID(context context.Context, id string) (*models.Product, error)
	GetProductsByBrandID(context context.Context, brandID string, limit, page int) ([]models.Product, error)
}

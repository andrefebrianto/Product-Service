package queries

import (
	"context"

	models "github.com/andrefebrianto/rest-api/bin/models"
)

// ProductQueries represent query function for product domain
type ProductQueries interface {
	Get(context context.Context) ([]models.Product, error)
	GetByID(context context.Context, id string) (models.Product, error)
	GetByBrandID(context context.Context, brandID string) (models.Product, error)
}

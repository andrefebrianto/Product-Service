package commands

import (
	"context"

	models "github.com/andrefebrianto/rest-api/bin/models"
)

// ProductQueries represent command function for product domain
type ProductQueries interface {
	Create(context context.Context, product *models.Product) ([]models.Product, error)
	Update(context context.Context, product *models.Product) (models.Product, error)
	Delete(context context.Context, id string) (models.Product, error)
}

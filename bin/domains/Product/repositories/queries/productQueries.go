package queries

import (
	"context"
	models "github.com/andrefebrianto/rest-api/bin/models"
)

type ProductQueries struct {
	Get(context context.Context) ([]models.Product, string, error)
	GetByID(id string) (Product, error)
}

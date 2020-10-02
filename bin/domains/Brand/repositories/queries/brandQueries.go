package queries

import (
	"context"

	models "github.com/andrefebrianto/rest-api/bin/models"
)

//BrandQueries represent query function for brand domain
type BrandQueries interface {
	Get(context context.Context) ([]models.Brand, error)
	GetByID(context context.Context, id string) (models.Brand, error)
}

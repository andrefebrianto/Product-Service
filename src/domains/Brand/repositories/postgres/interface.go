package repositories

import (
	"context"

	"github.com/andrefebrianto/rest-api/src/models"
)

//BrandCommands represent query function for brand domain
type BrandCommands interface {
	CreateBrand(context context.Context, brand *models.Brand) (*models.Brand, error)
	UpdateBrand(context context.Context, product *models.Brand) (*models.Brand, error)
	DeleteBrand(context context.Context, id string) error
}

//BrandQueries represent query function for brand domain
type BrandQueries interface {
	GetBrands(context context.Context, limit, page int) ([]models.Brand, error)
	GetBrandByID(context context.Context, id string) (models.Brand, error)
}

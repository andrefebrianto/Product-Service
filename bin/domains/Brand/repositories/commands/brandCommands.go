package commands

import (
	"context"

	"github.com/andrefebrianto/rest-api/bin/models"
)

//BrandCommands represent query function for brand domain
type BrandCommands interface {
	Create(context context.Context, brand *models.Brand) ([]models.Brand, error)
	Update(context context.Context, product *models.Brand) (models.Brand, error)
	Delete(context context.Context, id string) (models.Brand, error)
}

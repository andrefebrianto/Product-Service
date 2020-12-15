package queries

import (
	"context"
	"errors"

	models "github.com/andrefebrianto/rest-api/src/models"
)

//BrandQuery ...
type BrandQuery struct {
	dbConnection interface{}
}

//CreateRepository ...
func CreateRepository(connection interface{}) *BrandQuery {
	return &BrandQuery{connection}
}

//GetBrands ...
func (query BrandQuery) GetBrands(context context.Context, limit, page int) ([]models.Brand, error) {
	return nil, errors.New("Brand(s) not found")
}

//GetBrandByID ...
func (query BrandQuery) GetBrandByID(context context.Context, id string) (*models.Brand, error) {
	return nil, errors.New("Brand not found")
}

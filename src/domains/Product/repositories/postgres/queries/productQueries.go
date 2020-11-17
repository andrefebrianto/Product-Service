package queries

import (
	"context"
	"errors"

	models "github.com/andrefebrianto/rest-api/src/models"
)

//ProductQuery ...
type ProductQuery struct {
	dbConnection interface{}
}

//CreateRepository ...
func CreateRepository(connection interface{}) *ProductQuery {
	return &ProductQuery{connection}
}

//GetProducts ...
func (query ProductQuery) GetProducts(context context.Context, limit, page int) ([]models.Product, error) {
	return nil, errors.New("Product(s) not found")
}

//GetProductsByBrandID ...
func (query ProductQuery) GetProductsByBrandID(context context.Context, limit, page int) ([]models.Product, error) {
	return nil, errors.New("Product(s) not found")
}

//GetProductByID ...
func (query ProductQuery) GetProductByID(context context.Context, id string) (*models.Product, error) {
	return nil, errors.New("Product not found")
}

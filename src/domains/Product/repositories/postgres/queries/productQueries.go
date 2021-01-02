package queries

import (
	"context"
	"errors"

	models "github.com/andrefebrianto/rest-api/src/models"
	"github.com/go-pg/pg/v10"
)

type productQuery struct {
	dbConnection *pg.DB
}

//CreateRepository ...
func CreateRepository(connection *pg.DB) productQuery {
	return productQuery{connection}
}

//GetProducts ...
func (query productQuery) GetProducts(context context.Context, limit, page int) ([]models.Product, error) {
	var products []models.Product
	skip := (page - 1) * limit

	err := query.dbConnection.ModelContext(context, &products).Order("created_at DESC").Offset(skip).Limit(limit).Select()

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("Product(s) not found")
	}

	return products, nil
}

//GetProductsByBrandID ...
func (query productQuery) GetProductsByBrandID(context context.Context, brandID string, limit, page int) ([]models.Product, error) {
	var products []models.Product
	skip := (page - 1) * limit

	err := query.dbConnection.ModelContext(context, &products).Where("BrandID", brandID).Order("created_at DESC").Offset(skip).Limit(limit).Select()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("Product(s) not found")
	}

	return products, nil
}

//GetProductByID ...
func (query productQuery) GetProductByID(context context.Context, id string) (*models.Product, error) {
	product := new(models.Product)

	err := query.dbConnection.ModelContext(context, product).Where("id = ?", id).Select()

	if err != nil && err.Error() == "pg: no rows in result set" {
		return nil, errors.New("Product not found")
	} else if err != nil {
		return nil, err
	}

	return product, nil
}

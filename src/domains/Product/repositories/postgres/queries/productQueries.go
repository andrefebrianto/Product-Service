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

	err := query.dbConnection.ModelContext(context, products).Order("CreatedAt DESC").Offset(skip).Limit(limit).Select()
	if err == nil {
		return products, nil
	}
	return nil, errors.New("Product(s) not found")
}

//GetProductsByBrandID ...
func (query productQuery) GetProductsByBrandID(context context.Context, limit, page int, brandID string) ([]models.Product, error) {
	var products []models.Product
	skip := (page - 1) * limit

	err := query.dbConnection.ModelContext(context, products).Where("BrandID", brandID).Order("CreatedAt DESC").Offset(skip).Limit(limit).Select()
	if err == nil {
		return products, nil
	}
	return nil, errors.New("Product(s) not found")
}

//GetProductByID ...
func (query productQuery) GetProductByID(context context.Context, id string) (*models.Product, error) {
	product := new(models.Product)

	err := query.dbConnection.ModelContext(context, product).Where("ID", id).Order("CreatedAt DESC").Select()
	if err == nil {
		return product, nil
	}
	return nil, errors.New("Product not found")
}

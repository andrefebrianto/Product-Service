package queries

import (
	"context"
	"errors"

	models "github.com/andrefebrianto/rest-api/src/models"
	"github.com/go-pg/pg/v10"
)

//BrandQuery ...
type brandQuery struct {
	dbConnection *pg.DB
}

//CreateRepository ...
func CreateRepository(connection *pg.DB) brandQuery {
	return brandQuery{connection}
}

//GetBrands ...
func (query brandQuery) GetBrands(context context.Context, limit, page int) ([]models.Brand, error) {
	var brands []models.Brand
	skip := (page - 1) * limit

	err := query.dbConnection.ModelContext(context, brands).Order("CreatedAt DESC").Offset(skip).Limit(limit).Select()
	if err == nil {
		return brands, nil
	}
	return nil, errors.New("Brand(s) not found")
}

//GetBrandByID ...
func (query brandQuery) GetBrandByID(context context.Context, id string) (*models.Brand, error) {
	brand := new(models.Brand)

	err := query.dbConnection.ModelContext(context, brand).Where("ID", id).Order("CreatedAt DESC").Select()
	if err == nil {
		return brand, nil
	}
	return nil, errors.New("Brand not found")
}

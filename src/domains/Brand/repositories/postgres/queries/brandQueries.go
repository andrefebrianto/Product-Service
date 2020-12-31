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

	err := query.dbConnection.ModelContext(context, &brands).Order("created_at DESC").Offset(skip).Limit(limit).Select()
	if err != nil {
		return nil, err
	}

	if len(brands) == 0 {
		return nil, errors.New("Brand(s) not found")
	}

	return brands, nil
}

//GetBrandByID ...
func (query brandQuery) GetBrandByID(context context.Context, id string) (*models.Brand, error) {
	brand := new(models.Brand)

	err := query.dbConnection.ModelContext(context, brand).Where("id = ?", id).Select()
	if err.Error() == "pg: no rows in result set" {
		return nil, errors.New("Brand not found")
	}

	if err != nil {
		return nil, err
	}

	return brand, nil
}

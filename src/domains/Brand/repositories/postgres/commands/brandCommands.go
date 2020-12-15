package commands

import (
	"context"
	"errors"

	"github.com/andrefebrianto/rest-api/src/models"
)

//BrandCommand ...
type BrandCommand struct {
	dbConnection interface{}
}

//CreateRepository ...
func CreateRepository(connection interface{}) *BrandCommand {
	return &BrandCommand{connection}
}

//CreateBrand ...
func (command BrandCommand) CreateBrand(context context.Context, brand *models.Brand) (*models.Brand, error) {
	// command.dbConnection
	return nil, errors.New("Failed to create produt")
}

//UpdateBrand ...
func (command BrandCommand) UpdateBrand(context context.Context, brand *models.Brand) (*models.Brand, error) {
	return nil, errors.New("Failed to update produt")
}

//DeleteBrand ...
func (command BrandCommand) DeleteBrand(context context.Context, id string) error {
	return errors.New("Failed to delete produt")
}

package commands

import (
	"context"
	"errors"

	models "github.com/andrefebrianto/rest-api/src/models"
)

//ProductCommand ...
type ProductCommand struct {
	dbConnection interface{}
}

//CreateRepository ...
func CreateRepository(connection interface{}) *ProductCommand {
	return &ProductCommand{connection}
}

//CreateProduct ...
func (command ProductCommand) CreateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	// command.dbConnection
	return nil, errors.New("Failed to create produt")
}

//UpdateProduct ...
func (command ProductCommand) UpdateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	return nil, errors.New("Failed to update produt")
}

//DeleteProduct ...
func (command ProductCommand) DeleteProduct(context context.Context, id string) error {
	return errors.New("Failed to delete produt")
}

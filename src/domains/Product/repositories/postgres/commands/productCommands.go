package commands

import (
	"context"
	"errors"

	models "github.com/andrefebrianto/rest-api/src/models"
	"github.com/go-pg/pg/v10"
)

//ProductCommand ...
type productCommand struct {
	dbConnection *pg.DB
}

//CreateRepository ...
func CreateRepository(connection *pg.DB) productCommand {
	return productCommand{connection}
}

//CreateProduct ...
func (command productCommand) CreateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, product).Insert()
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		return product, nil
	}

	return nil, errors.New("Failed to create produt")
}

//UpdateProduct ...
func (command productCommand) UpdateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, product).WherePK().Update()
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		return product, nil
	}

	return nil, errors.New("Failed to update produt")
}

//UpdateProductStock ...
func (command productCommand) UpdateProductStock(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, product).Column("Stock", "UpdatedAt").WherePK().Update()
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		return product, nil
	}

	return nil, errors.New("Failed to update produt")
}

//DeleteProduct ...
func (command productCommand) DeleteProduct(context context.Context, id string) error {
	product := &models.Product{
		ID: id,
	}
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, product).WherePK().Update()
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		return nil
	}

	return errors.New("Failed to delete produt")
}

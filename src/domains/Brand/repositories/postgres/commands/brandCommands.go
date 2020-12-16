package commands

import (
	"context"
	"errors"

	"github.com/andrefebrianto/rest-api/src/models"
	"github.com/go-pg/pg/v10"
)

//BrandCommand ...
type BrandCommand struct {
	dbConnection *pg.DB
}

//CreateRepository ...
func CreateRepository(connection *pg.DB) *BrandCommand {
	return &BrandCommand{connection}
}

//CreateBrand ...
func (command BrandCommand) CreateBrand(context context.Context, brand *models.Brand) (*models.Brand, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, brand).Insert()
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		return brand, nil
	}

	return nil, errors.New("Failed to create produt")
}

//UpdateBrand ...
func (command BrandCommand) UpdateBrand(context context.Context, brand *models.Brand) (*models.Brand, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, brand).WherePK().Update()
		if err != nil {
			return err
		}

		return err
	})

	if err == nil {
		return brand, nil
	}

	return nil, errors.New("Failed to update produt")
}

//DeleteBrand ...
func (command BrandCommand) DeleteBrand(context context.Context, id string) error {
	brand := &models.Brand{
		ID: id,
	}
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, brand).WherePK().Update()
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

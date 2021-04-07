package commands

import (
	"context"

	"github.com/andrefebrianto/rest-api/src/models"
	"github.com/go-pg/pg/v10"
)

//BrandCommand ...
type brandCommand struct {
	pgClient *pg.DB
}

//CreateRepository ...
func CreateRepository(pgClient *pg.DB) brandCommand {
	return brandCommand{pgClient}
}

//CreateBrand ...
func (command brandCommand) CreateBrand(context context.Context, brand *models.Brand) (*models.Brand, error) {
	err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, brand).Insert()
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return brand, nil
}

//UpdateBrand ...
func (command brandCommand) UpdateBrand(context context.Context, brand *models.Brand) (*models.Brand, error) {
	err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, brand).Column("name", "updated_at").WherePK().Update()
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return brand, nil
}

//DeleteBrand ...
func (command brandCommand) DeleteBrand(context context.Context, id string) error {
	brand := &models.Brand{
		ID: id,
	}
	err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, brand).WherePK().Delete()
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return err
	}

	return nil
}

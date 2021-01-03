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

	if err != nil {
		return nil, err
	}

	return product, nil
}

//UpdateProduct ...
func (command productCommand) UpdateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, product).Column("name", "price", "brand_id", "description", "stock", "sold", "updated_at").WherePK().Update()
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}

//UpdateProductStock ...
func (command productCommand) UpdateProductStock(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		result, err := dbTransaction.ModelContext(context, product).Column("stock", "updated_at").WherePK().Update()

		err = dbTransaction.ModelContext(context, product).WherePK().Select()

		if err != nil {
			return err
		}

		if result.RowsAffected() == 0 {
			return errors.New("Product not found")
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}

//DeleteProduct ...
func (command productCommand) DeleteProduct(context context.Context, id string) error {
	product := &models.Product{
		ID: id,
	}
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		_, err := dbTransaction.ModelContext(context, product).WherePK().Delete()
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

//UpdateProductStock ...
func (command productCommand) UpdatePurchasedStock(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.dbConnection.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		currentProduct := new(models.Product)
		err := dbTransaction.ModelContext(context, currentProduct).Where("id = ?", product.ID).For("UPDATE").Select()
		stock, sold, err := currentProduct.BuyProduct(product.Stock)
		if err != nil {
			return err
		}

		product.Stock = stock
		product.Sold = sold
		result, err := dbTransaction.ModelContext(context, product).Column("stock", "sold", "updated_at").WherePK().Update()

		if err != nil {
			return err
		}

		if result.RowsAffected() == 0 {
			return errors.New("Product not found")
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}

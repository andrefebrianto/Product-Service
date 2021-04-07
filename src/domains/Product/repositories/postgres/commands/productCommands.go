package commands

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	models "github.com/andrefebrianto/rest-api/src/models"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-pg/pg/v10"
)

var ELASTIC_PRODUCT_CATALOG_INDEX = "product-catalog"

//ProductCommand ...
type productCommand struct {
	pgClient *pg.DB
	esClient *elasticsearch.Client
}

type productIndex struct {
	ID        string
	Name      string
	Price     int
	BrandName string
	Stock     int
	Sold      int
}

//CreateRepository ...
func CreateRepository(pgClient *pg.DB, esClient *elasticsearch.Client) productCommand {
	return productCommand{pgClient, esClient}
}

//CreateProduct ...
func (command productCommand) CreateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	pgInsertErrorCh := make(chan error)
	pgSelectErrorCh := make(chan error)
	esErrorCh := make(chan error)
	brandCh := make(chan *models.Brand)

	go func() {
		err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
			_, err := dbTransaction.ModelContext(context, product).Insert()
			if err != nil {
				return err
			}

			return err
		})
		pgInsertErrorCh <- err
	}()

	go func() {
		brand := new(models.Brand)
		err := command.pgClient.ModelContext(context, brand).Where("id = ?", product.BrandId).Select()
		pgSelectErrorCh <- err
		brandCh <- brand
	}()

	pgErr := <-pgInsertErrorCh
	if pgErr != nil {
		return nil, pgErr
	}

	pgErr = <-pgSelectErrorCh
	if pgErr != nil {
		return nil, pgErr
	}

	brand := <-brandCh

	go func() {
		productIdx := productIndex{ID: product.ID, Name: product.Name, BrandName: brand.Name, Price: product.Price, Stock: product.Stock, Sold: product.Sold}
		stringObject, _ := json.Marshal(productIdx)

		_, err := command.esClient.Index(ELASTIC_PRODUCT_CATALOG_INDEX, strings.NewReader(string(stringObject)), command.esClient.Index.WithDocumentID(product.ID))

		esErrorCh <- err
	}()

	esErr := <-esErrorCh

	if esErr != nil {
		return nil, esErr
	}

	return product, nil
}

//UpdateProduct ...
func (command productCommand) UpdateProduct(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
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
	err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		result, err := dbTransaction.ModelContext(context, product).Column("stock", "updated_at").WherePK().Update()
		if err != nil {
			return err
		}

		err = dbTransaction.ModelContext(context, product).WherePK().Select()

		if err != nil {
			return err
		}

		if result.RowsAffected() == 0 {
			return errors.New("product not found")
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
	pgErrCh := make(chan error)
	esErrCh := make(chan error)

	go func() {
		product := &models.Product{
			ID: id,
		}
		err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
			_, err := dbTransaction.ModelContext(context, product).WherePK().Delete()
			if err != nil {
				return err
			}

			return err
		})
		pgErrCh <- err
	}()

	go func() {
		_, err := command.esClient.Delete(ELASTIC_PRODUCT_CATALOG_INDEX, id)
		esErrCh <- err
	}()

	pgErr := <-pgErrCh
	if pgErr != nil {
		return pgErr
	}
	esErr := <-esErrCh
	if pgErr != nil {
		return esErr
	}

	return nil
}

//UpdateProductStock ...
func (command productCommand) UpdatePurchasedStock(context context.Context, product *models.Product) (*models.Product, error) {
	err := command.pgClient.RunInTransaction(context, func(dbTransaction *pg.Tx) error {
		currentProduct := new(models.Product)
		err := dbTransaction.ModelContext(context, currentProduct).Where("id = ?", product.ID).For("UPDATE").Select()
		if err != nil {
			return err
		}
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
			return errors.New("product not found")
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}

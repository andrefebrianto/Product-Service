package usecases

import (
	"context"
	"time"

	brandrepo "github.com/andrefebrianto/rest-api/src/domains/Product/repositories/postgres"
	"github.com/andrefebrianto/rest-api/src/models"
)

type ProductUseCase struct {
	commandRepository brandrepo.ProductCommand
	queryRepository   brandrepo.ProductQueries
	contextTimeout    time.Duration
}

func CreateProductUseCase(command brandrepo.ProductCommand, query brandrepo.ProductQueries, timeout time.Duration) ProductUseCase {
	return ProductUseCase{
		commandRepository: command,
		queryRepository:   query,
		contextTimeout:    timeout,
	}
}

func (useCase *ProductUseCase) CreateProduct(mainContext context.Context, product *models.Product) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	createdProduct, err := useCase.commandRepository.CreateProduct(contextWithTimeout, product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func (useCase *ProductUseCase) UpdateProduct(mainContext context.Context, product *models.Product) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	updatedProduct, err := useCase.commandRepository.UpdateProduct(contextWithTimeout, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (useCase *ProductUseCase) UpdateProductStock(mainContext context.Context, id string, stock int) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	product := &models.Product{ID: id, Stock: stock, UpdatedAt: time.Now()}

	updatedProduct, err := useCase.commandRepository.UpdateProductStock(contextWithTimeout, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (useCase *ProductUseCase) PurchaseProduct(mainContext context.Context, id string, stock int) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	product := &models.Product{ID: id, Stock: stock, UpdatedAt: time.Now()}

	updatedProduct, err := useCase.commandRepository.UpdatePurchasedStock(contextWithTimeout, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (useCase *ProductUseCase) DeleteProduct(mainContext context.Context, id string) error {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	err := useCase.commandRepository.DeleteProduct(contextWithTimeout, id)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *ProductUseCase) GetProducts(mainContext context.Context, limit, page int) ([]models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	products, err := useCase.queryRepository.GetProducts(contextWithTimeout, limit, page)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (useCase *ProductUseCase) GetProductByID(mainContext context.Context, id string) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	product, err := useCase.queryRepository.GetProductByID(contextWithTimeout, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (useCase *ProductUseCase) GetProductByBrandID(mainContext context.Context, brandID string, limit, page int) ([]models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	products, err := useCase.queryRepository.GetProductsByBrandID(contextWithTimeout, brandID, limit, page)
	if err != nil {
		return nil, err
	}
	return products, nil
}

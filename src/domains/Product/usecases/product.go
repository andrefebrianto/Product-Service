package usecases

import (
	"context"
	"time"

	"github.com/andrefebrianto/rest-api/src/domains/Product/repositories/postgres"
	"github.com/andrefebrianto/rest-api/src/models"
)

type productUseCase struct {
	commandRepository postgres.ProductCommand
	queryRepository   postgres.ProductQueries
	contextTimeout    time.Duration
}

func CreateProductUseCase() *productUseCase {
	return &productUseCase{}
}

func (useCase *productUseCase) CreateProduct(mainContext context.Context, product *models.Product) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	createdProduct, err := useCase.commandRepository.CreateProduct(contextWithTimeout, product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func (useCase *productUseCase) UpdateProduct(mainContext context.Context, product *models.Product) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	updatedProduct, err := useCase.commandRepository.UpdateProduct(contextWithTimeout, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (useCase *productUseCase) DeleteProduct(mainContext context.Context, id string) error {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	err := useCase.commandRepository.DeleteProduct(contextWithTimeout, id)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *productUseCase) GetProducts(mainContext context.Context, limit, page int) ([]models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	products, err := useCase.queryRepository.GetProducts(contextWithTimeout, limit, page)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (useCase *productUseCase) GetProductByID(mainContext context.Context, id string) (*models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	product, err := useCase.queryRepository.GetProductByID(contextWithTimeout, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (useCase *productUseCase) GetProductByBrandID(mainContext context.Context, brandID string, limit, page int) ([]models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	products, err := useCase.queryRepository.GetProductsByBrandID(contextWithTimeout, brandID, limit, page)
	if err != nil {
		return nil, err
	}
	return products, nil
}

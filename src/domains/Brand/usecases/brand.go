package usecases

import (
	"context"
	"time"

	brandrepo "github.com/andrefebrianto/rest-api/src/domains/Brand/repositories/postgres"
	"github.com/andrefebrianto/rest-api/src/models"
)

type BrandUseCase struct {
	commandRepository brandrepo.BrandCommands
	queryRepository   brandrepo.BrandQueries
	contextTimeout    time.Duration
}

func CreateProductUseCase() *BrandUseCase {
	return &BrandUseCase{}
}

func (useCase *BrandUseCase) CreateBrand(mainContext context.Context, brand *models.Brand) (*models.Brand, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	createdBrand, err := useCase.commandRepository.CreateBrand(contextWithTimeout, brand)
	if err != nil {
		return nil, err
	}
	return createdBrand, nil
}

func (useCase *BrandUseCase) UpdateProduct(mainContext context.Context, brand *models.Brand) (*models.Brand, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	updatedProduct, err := useCase.commandRepository.UpdateBrand(contextWithTimeout, brand)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (useCase *BrandUseCase) DeleteProduct(mainContext context.Context, id string) error {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	err := useCase.commandRepository.DeleteBrand(contextWithTimeout, id)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *BrandUseCase) GetProducts(mainContext context.Context, limit, page int) ([]models.Product, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	products, err := useCase.GetProducts(contextWithTimeout, limit, page)
	if err != nil {
		return nil, err
	}
	return products, nil
}

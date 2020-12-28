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

func CreateBrandUseCase(command brandrepo.BrandCommands, query brandrepo.BrandQueries, timeout time.Duration) BrandUseCase {
	return BrandUseCase{
		commandRepository: command,
		queryRepository:   query,
		contextTimeout:    timeout,
	}
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

func (useCase *BrandUseCase) UpdateBrand(mainContext context.Context, brand *models.Brand) (*models.Brand, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	updatedBrand, err := useCase.commandRepository.UpdateBrand(contextWithTimeout, brand)
	if err != nil {
		return nil, err
	}
	return updatedBrand, nil
}

func (useCase *BrandUseCase) DeleteBrand(mainContext context.Context, id string) error {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	err := useCase.commandRepository.DeleteBrand(contextWithTimeout, id)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *BrandUseCase) GetBrands(mainContext context.Context, limit, page int) ([]models.Brand, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	brands, err := useCase.queryRepository.GetBrands(contextWithTimeout, limit, page)
	if err != nil {
		return nil, err
	}
	return brands, nil
}

func (useCase *BrandUseCase) GetBrandByID(mainContext context.Context, id string) (*models.Brand, error) {
	contextWithTimeout, cancel := context.WithTimeout(mainContext, useCase.contextTimeout)
	defer cancel()

	brand, err := useCase.queryRepository.GetBrandByID(contextWithTimeout, id)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

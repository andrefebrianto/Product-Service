package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// BrandHandler ...
type BrandHandler struct {
	useCase string
}

// GetBrandByID ...
func (handler *BrandHandler) GetBrandByID(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// GetBrands ...
func (handler *BrandHandler) GetBrands(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// AddBrand ...
func (handler *BrandHandler) AddBrand(context echo.Context) error {

	return context.JSON(http.StatusCreated, nil)
}

// DeleteBrand ...
func (handler *BrandHandler) DeleteBrand(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// UpdateBrand ...
func (handler *BrandHandler) UpdateBrand(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

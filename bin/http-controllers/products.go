package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// ProductHandler ...
type ProductHandler struct {
	useCase string
}

// GetProductByID ...
func (handler *ProductHandler) GetProductByID(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// GetProducts ...
func (handler *ProductHandler) GetProducts(context echo.Context) error {
	return context.JSON(http.StatusOK, nil)
}

// PurchaseProduct ...
func (handler *ProductHandler) PurchaseProduct(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// RestockProduct ...
func (handler *ProductHandler) RestockProduct(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// UpdateProduct ...
func (handler *ProductHandler) UpdateProduct(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// AddProduct ...
func (handler *ProductHandler) AddProduct(context echo.Context) error {

	return context.JSON(http.StatusCreated, nil)
}

// DeleteProduct ...
func (handler *ProductHandler) DeleteProduct(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

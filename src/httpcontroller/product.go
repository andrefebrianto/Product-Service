package httpcontroller

import (
	"net/http"
	"strconv"

	usecase "github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	"github.com/andrefebrianto/rest-api/src/models"
	"github.com/labstack/echo"
)

// ProductHandler ...
type ProductHandler struct {
	useCase usecase.ProductUseCase
}

// GetProductByID ...
func (handler *ProductHandler) GetProductByID(context echo.Context) error {
	id := context.QueryParam("id")
	ctx := context.Request().Context()

	product, err := handler.useCase.GetProductByID(ctx, id)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if product == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Product not found"})
	}

	return context.JSON(http.StatusOK, product)
}

// GetProducts ...
func (handler *ProductHandler) GetProducts(context echo.Context) error {
	page, _ := strconv.Atoi(context.QueryParam("page"))
	limit, _ := strconv.Atoi(context.QueryParam("limit"))

	ctx := context.Request().Context()

	products, err := handler.useCase.GetProducts(ctx, limit, page)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if products == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Product(s) not found"})
	}

	return context.JSON(http.StatusOK, products)
}

// PurchaseProduct ...
func (handler *ProductHandler) PurchaseProduct(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// RestockProduct ...
func (handler *ProductHandler) UpdateProductStock(context echo.Context) error {
	id := context.QueryParam("id")
	ctx := context.Request().Context()
	payload := map[string]interface{}{}

	if err := context.Bind(&payload); err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	product, err := handler.useCase.UpdateProductStock(ctx, id, payload["stock"].(int))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if product == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Product not found"})
	}

	return context.JSON(http.StatusOK, nil)
}

// UpdateProduct ...
func (handler *ProductHandler) UpdateProduct(context echo.Context) error {
	var product *models.Product
	err := context.Bind(product)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := context.Request().Context()

	updatedProduct, err := handler.useCase.UpdateProduct(ctx, product)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, updatedProduct)
}

// AddProduct ...
func (handler *ProductHandler) AddProduct(context echo.Context) error {
	var product *models.Product
	err := context.Bind(product)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := context.Request().Context()

	createdProduct, err := handler.useCase.CreateProduct(ctx, product)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, createdProduct)
}

// DeleteProduct ...
func (handler *ProductHandler) DeleteProduct(context echo.Context) error {
	id := context.Param("id")

	ctx := context.Request().Context()

	err := handler.useCase.DeleteProduct(ctx, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, nil)
}

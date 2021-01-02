package httpcontroller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	usecase "github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	"github.com/andrefebrianto/rest-api/src/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ProductHandler ...
type ProductHandler struct {
	UseCase usecase.ProductUseCase
}

type productPayloadBuffer struct {
	Name        string
	Price       int
	BrandId     string
	Description string
	Stock       int
	Sold        int
}

func CreateProductHandler(e *echo.Echo, usecase usecases.ProductUseCase) {
	handler := &ProductHandler{UseCase: usecase}

	e.POST("/api/products", handler.AddProduct)
	e.DELETE("/api/products/:id", handler.DeleteProduct)
	e.PUT("/api/products/:id", handler.UpdateProduct)
	e.PATCH("api/products/:id/stocks", handler.UpdateProductStock)
	e.GET("/api/products", handler.GetProducts)
	e.GET("/api/products/:id", handler.GetProductByID)
}

// GetProductByID ...
func (handler *ProductHandler) GetProductByID(context echo.Context) error {
	id := context.Param("id")
	ctx := context.Request().Context()

	product, err := handler.UseCase.GetProductByID(ctx, id)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Product retrieved", Data: product})
}

// GetProducts ...
func (handler *ProductHandler) GetProducts(context echo.Context) error {
	page, _ := strconv.Atoi(context.QueryParam("page"))
	limit, _ := strconv.Atoi(context.QueryParam("limit"))

	ctx := context.Request().Context()

	products, err := handler.UseCase.GetProducts(ctx, limit, page)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if products == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Product(s) not found"})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Product(s) retrieved", Data: products})
}

// PurchaseProduct ...
func (handler *ProductHandler) PurchaseProduct(context echo.Context) error {

	return context.JSON(http.StatusOK, nil)
}

// RestockProduct ...
func (handler *ProductHandler) UpdateProductStock(context echo.Context) error {
	ctx := context.Request().Context()
	payload := map[string]interface{}{}

	if err := context.Bind(&payload); err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	product, err := handler.UseCase.UpdateProductStock(ctx, payload["id"].(string), int(payload["stock"].(float64)))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if product == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Product not found"})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Stock updated", Data: product})
}

// UpdateProduct ...
func (handler *ProductHandler) UpdateProduct(context echo.Context) error {
	var product models.Product
	err := context.Bind(&product)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	ctx := context.Request().Context()
	product.UpdatedAt = time.Now()

	updatedProduct, err := handler.UseCase.UpdateProduct(ctx, &product)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Product updated", Data: updatedProduct})
}

// AddProduct ...
func (handler *ProductHandler) AddProduct(context echo.Context) error {
	// var buffer productPayloadBuffer
	var product models.Product
	err := context.Bind(&product)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	generatedId, err := uuid.NewRandom()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	product.ID = generatedId.String()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	ctx := context.Request().Context()

	createdProduct, err := handler.UseCase.CreateProduct(ctx, &product)
	if err != nil && err.Error() == "ERROR #23503 insert or update on table \"products\" violates foreign key constraint \"products_brand_id_fkey\"" {
		return context.JSON(http.StatusForbidden, ResponseError{Message: "Brand not found"})
	} else if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Product created", Data: createdProduct})
}

// DeleteProduct ...
func (handler *ProductHandler) DeleteProduct(context echo.Context) error {
	id := context.Param("id")

	ctx := context.Request().Context()

	err := handler.UseCase.DeleteProduct(ctx, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Product deleted"})
}

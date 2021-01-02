package httpcontroller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/andrefebrianto/rest-api/src/domains/Brand/usecases"
	brandusecase "github.com/andrefebrianto/rest-api/src/domains/Brand/usecases"
	"github.com/andrefebrianto/rest-api/src/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
	Data    interface{}
}

// BrandHandler ...
type BrandHandler struct {
	UseCase brandusecase.BrandUseCase
}

func CreateBrandHandler(e *echo.Echo, usecase usecases.BrandUseCase) {
	handler := &BrandHandler{UseCase: usecase}

	e.POST("/api/brands", handler.AddBrand)
	e.DELETE("/api/brands/:id", handler.DeleteBrand)
	e.PUT("/api/brands/:id", handler.UpdateBrand)
	e.GET("/api/brands", handler.GetBrands)
	e.GET("/api/brands/:id", handler.GetBrandByID)
}

// GetBrandByID ...
func (handler *BrandHandler) GetBrandByID(context echo.Context) error {
	id := context.Param("id")
	ctx := context.Request().Context()

	brand, err := handler.UseCase.GetBrandByID(ctx, id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Brand retrieved", Data: brand})
}

// GetBrands ...
func (handler *BrandHandler) GetBrands(context echo.Context) error {
	page, _ := strconv.Atoi(context.QueryParam("page"))
	limit, _ := strconv.Atoi(context.QueryParam("limit"))

	ctx := context.Request().Context()

	brands, err := handler.UseCase.GetBrands(ctx, limit, page)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if brands == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Brand(s) not found"})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Brand(s) retrieved", Data: brands})
}

// AddBrand ...
func (handler *BrandHandler) AddBrand(context echo.Context) error {
	var brand models.Brand
	err := context.Bind(&brand)

	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	ctx := context.Request().Context()

	generatedId, err := uuid.NewRandom()

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	brand.ID = generatedId.String()
	brand.CreatedAt = time.Now()
	brand.UpdatedAt = time.Now()

	createdBrand, err := handler.UseCase.CreateBrand(ctx, &brand)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusCreated, ResponseError{Message: "Brand created", Data: createdBrand})
}

// DeleteBrand ...
func (handler *BrandHandler) DeleteBrand(context echo.Context) error {
	id := context.Param("id")

	ctx := context.Request().Context()

	err := handler.UseCase.DeleteBrand(ctx, id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Brand deleted"})
}

func (handler *BrandHandler) UpdateBrand(context echo.Context) error {
	var brand models.Brand
	err := context.Bind(&brand)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}

	ctx := context.Request().Context()
	brand.UpdatedAt = time.Now()

	updatedBrand, err := handler.UseCase.UpdateBrand(ctx, &brand)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Brand updated", Data: updatedBrand})
}

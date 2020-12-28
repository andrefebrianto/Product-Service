package httpcontroller

import (
	"net/http"
	"strconv"

	brandusecase "github.com/andrefebrianto/rest-api/src/domains/Brand/usecases"
	"github.com/andrefebrianto/rest-api/src/models"
	"github.com/labstack/echo"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// BrandHandler ...
type BrandHandler struct {
	useCase brandusecase.BrandUseCase
}

// GetBrandByID ...
func (handler *BrandHandler) GetBrandByID(context echo.Context) error {
	id := context.QueryParam("id")
	ctx := context.Request().Context()

	brand, err := handler.useCase.GetBrandByID(ctx, id)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if brand == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Brand not found"})
	}

	return context.JSON(http.StatusOK, brand)
}

// GetBrands ...
func (handler *BrandHandler) GetBrands(context echo.Context) error {
	page, _ := strconv.Atoi(context.QueryParam("page"))
	limit, _ := strconv.Atoi(context.QueryParam("limit"))

	ctx := context.Request().Context()

	brands, err := handler.useCase.GetBrands(ctx, limit, page)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	if brands == nil {
		return context.JSON(http.StatusNotFound, ResponseError{Message: "Brand(s) not found"})
	}

	return context.JSON(http.StatusOK, brands)
}

// AddBrand ...
func (handler *BrandHandler) AddBrand(context echo.Context) error {
	var brand models.Brand
	err := context.Bind(brand)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := context.Request().Context()

	createdBrand, err := handler.useCase.CreateBrand(ctx, &brand)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, createdBrand)
}

// DeleteBrand ...
func (handler *BrandHandler) DeleteBrand(context echo.Context) error {
	id := context.Param("id")

	ctx := context.Request().Context()

	err := handler.useCase.DeleteBrand(ctx, id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, nil)
}

func (handler *BrandHandler) UpdateBrand(context echo.Context) error {
	var brand models.Brand
	err := context.Bind(brand)
	if err != nil {
		return context.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := context.Request().Context()

	updatedBrand, err := handler.useCase.UpdateBrand(ctx, &brand)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, updatedBrand)
}

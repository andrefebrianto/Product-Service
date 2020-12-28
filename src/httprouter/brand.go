package httprouter

import (
	"github.com/andrefebrianto/rest-api/src/domains/Brand/usecases"
	controller "github.com/andrefebrianto/rest-api/src/httpcontroller"
	"github.com/labstack/echo/v4"
)

func CreateBrandHttpRouter(e *echo.Echo, usecase usecases.BrandUseCase) {
	handler := &controller.BrandHandler{UseCase: usecase}

	e.POST("/api/brands", handler.AddBrand)
	e.DELETE("/api/brands/:id", handler.DeleteBrand)
	e.GET("/api/brands", handler.GetBrands)
	e.GET("/api/brands/:id", handler.GetBrandByID)
}

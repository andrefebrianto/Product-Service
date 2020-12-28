package httprouter

import (
	controller "github.com/andrefebrianto/rest-api/src/httpcontroller"
	"github.com/labstack/echo/v4"
)

func CreateBrandHttpRouter(e *echo.Echo) {
	handler := &controller.BrandHandler{}

	e.POST("/api/brands", handler.AddBrand)
	e.DELETE("/api/brands/:id", handler.DeleteBrand)
	e.GET("/api/brands", handler.GetBrands)
	e.GET("/api/brands/:id", handler.GetBrandByID)
}

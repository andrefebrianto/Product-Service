package httprouter

import (
	"github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	controller "github.com/andrefebrianto/rest-api/src/httpcontroller"
	"github.com/labstack/echo/v4"
)

func CreateProductHttpRouter(e *echo.Echo, usecase usecases.ProductUseCase) {
	handler := &controller.ProductHandler{UseCase: usecase}

	e.POST("/api/products", handler.AddProduct)
	e.DELETE("/api/products/:id", handler.DeleteProduct)
	e.GET("/api/products", handler.GetProducts)
	e.GET("/api/products/:id", handler.GetProductByID)
}

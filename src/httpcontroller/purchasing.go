package httpcontroller

import (
	"net/http"

	"github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	usecase "github.com/andrefebrianto/rest-api/src/domains/Product/usecases"
	"github.com/labstack/echo/v4"
)

// ProductHandler ...
type PurchaseHandler struct {
	UseCase usecase.ProductUseCase
}

func CreatePurchaseHandler(e *echo.Echo, usecase usecases.ProductUseCase) {
	handler := &PurchaseHandler{UseCase: usecase}

	e.POST("/api/purchases", handler.PurchaseProduct)
}

// GetProductByID ...
func (handler *PurchaseHandler) PurchaseProduct(context echo.Context) error {
	payload := map[string]interface{}{}

	if err := context.Bind(&payload); err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	ctx := context.Request().Context()

	_, err := handler.UseCase.PurchaseProduct(ctx, payload["productId"].(string), int(payload["stock"].(float64)))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, ResponseError{Message: "Product purchased"})
}

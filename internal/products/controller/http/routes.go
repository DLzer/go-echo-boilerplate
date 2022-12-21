package http

import (
	"github.com/DLzer/go-echo-boilerplate/internal/products"
	"github.com/labstack/echo/v4"
)

// Map product routes
func MapProductsRoutes(productsGroup *echo.Group, h products.Handlers) {
	productsGroup.POST("/create", h.Create())
	productsGroup.PATCH("/update/:id", h.Update())
	productsGroup.DELETE("/:id", h.Delete())
	productsGroup.GET("/:id", h.GetByID())
	productsGroup.GET("", h.All())
}

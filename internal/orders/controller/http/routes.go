package http

import (
	"github.com/DLzer/go-echo-boilerplate/internal/orders"
	"github.com/labstack/echo/v4"
)

// Map order routes
func MapOrderRoutes(ordersGroup *echo.Group, h orders.Handlers) {
	ordersGroup.POST("/create", h.Create())
	ordersGroup.PATCH("/update/:id", h.Update())
	ordersGroup.DELETE("/:id", h.Delete())
	ordersGroup.GET("/:id", h.GetByID())
	ordersGroup.GET("", h.All())
}

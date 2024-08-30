package http

import (
	"github.com/DLzer/go-echo-boilerplate/internal/users"
	apiMiddlewares "github.com/DLzer/go-echo-boilerplate/pkg/middleware"
	"github.com/labstack/echo/v4"
)

// Map Routes
func UserRoutes(group *echo.Group, h users.Handler, mw apiMiddlewares.MiddlewareManager) {
	group.POST("", h.Create())
	group.PUT("/:id", h.Update())
	group.DELETE("/:id", h.Delete())
	group.GET("/:id", h.GetByID())
	group.GET("", h.GetList())
	group.GET("/search", h.Search())
}

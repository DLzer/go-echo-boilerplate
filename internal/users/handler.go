package users

import "github.com/labstack/echo/v4"

type Handler interface {
	// Create
	Create() echo.HandlerFunc
	// Update
	Update() echo.HandlerFunc
	// Delete
	Delete() echo.HandlerFunc
	// Get By ID
	GetByID() echo.HandlerFunc
	// Get List
	GetList() echo.HandlerFunc
	// Name Search
	Search() echo.HandlerFunc
}

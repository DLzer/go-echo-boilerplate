package http

import (
	"net/http"

	"github.com/DLzer/go-echo-boilerplate/internal/config"
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/users"
	"github.com/DLzer/go-echo-boilerplate/pkg/httpErrors"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// User Handler
type user_handler struct {
	cfg     *config.Config
	service users.Service
	logger  logger.Logger
}

// User Handler Constructor
func NewUsersHandler(cfg *config.Config, service users.Service, logger logger.Logger) users.Handler {
	return &user_handler{
		cfg:     cfg,
		service: service,
		logger:  logger,
	}
}

// Create godoc
// @Summary Creates a User
// @Description Accepts a user create model
// @Tags  Users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param payload body models.UserCreate false "User"
// @Router /users [post]
func (h user_handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := &models.UserCreate{}

		if err := c.Bind(r); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := h.service.Create(r)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// c.Response().Header().Set("Cache-Control", "public, max-age=300")
		return c.JSON(http.StatusOK, user)
	}
}

// Update godoc
// @Summary Update a user
// @Description Accepts a user to update, responds with the updated users
// @Tags  Users
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param payload body models.User false "A user"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (h user_handler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uuid := c.Param("id")

		r := &models.UserUpdate{}
		if err := c.Bind(r); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		parsed_user_uuid, err := uuid.Parse(user_uuid)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := h.service.Update(r, parsed_user_uuid)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// c.Response().Header().Set("Cache-Control", "public, max-age=300")
		return c.JSON(http.StatusOK, user)
	}
}

// Delete by ID
// @Summary Deletes a user
// @Description Accepts a user UUID as a path parameter, and returns a true or false body.
// @Tags  Users
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} boolean
// @Router /users/{id} [delete]
func (h user_handler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uuid := c.Param("id")

		parsed_user_uuid, err := uuid.Parse(user_uuid)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := h.service.Delete(parsed_user_uuid)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// c.Response().Header().Set("Cache-Control", "public, max-age=300")
		return c.JSON(http.StatusOK, user)
	}
}

// @Summary Get user by ID
// @Description Accepts a UUID as a query parameter, and returns the User
// @Tags  Users
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (h user_handler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_uuid := c.Param("id")

		parsed_user_uuid, err := uuid.Parse(user_uuid)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := h.service.GetByID(parsed_user_uuid)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// c.Response().Header().Set("Cache-Control", "public, max-age=300")
		return c.JSON(http.StatusOK, user)
	}
}

// Get User List godoc
// @Summary Get user list
// @Description Accepts pagination query parameters, and returns a list of Users.
// @Tags  Users
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200
// @Router /users [get]
func (h user_handler) GetList() echo.HandlerFunc {
	return func(c echo.Context) error {
		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		list, err := h.service.GetList(pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// c.Response().Header().Set("Cache-Control", "public, max-age=300")
		return c.JSON(http.StatusOK, list)
	}
}

// Search godoc
// @Summary Performs a TSVector Search
// @Description Accepts a user name as a query param, performs a tsvector search and returns a list of matches
// @Tags Users
// @Accept json
// @Produce json
// @Param name query string false "filter name" Format(name)
// @Router /users/search [get]
func (h user_handler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_name := c.QueryParam("name")

		users, err := h.service.Search(user_name)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		// c.Response().Header().Set("Cache-Control", "public, max-age=300")
		return c.JSON(http.StatusOK, users)
	}
}

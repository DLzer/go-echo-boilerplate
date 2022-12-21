package http

import (
	"net/http"

	"github.com/DLzer/go-echo-boilerplate/config"
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/products"
	"github.com/DLzer/go-echo-boilerplate/pkg/httpErrors"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

// Products Handler
type productsHandlers struct {
	cfg             *config.Config
	productsService products.Service
	logger          logger.Logger
}

// NewProductsHandlers Products handlers constructor
func NewProductsHandlers(cfg *config.Config, productsService products.Service, logger logger.Logger) products.Handlers {
	return &productsHandlers{cfg: cfg, productsService: productsService, logger: logger}
}

// Create godoc
// @Summary Create products
// @Description Create products handler
// @Tags Products
// @Accept json
// @Produce json
// @Success 201 {object} models.ProductRequest
// @Router /products/create [post]
func (h productsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "productsHandlers.Create")
		defer span.Finish()

		p := &models.ProductRequest{}
		if err := c.Bind(p); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		created, err := h.productsService.Create(ctx, p)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, created)
	}
}

// Create godoc
// @Summary Update products
// @Description Update products handler
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 201 {object} models.ProductRequest
// @Router /products/update/{id} [patch]
func (h productsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "productsHandlers.Update")
		defer span.Finish()

		id := c.Param("id")

		p := &models.ProductRequest{}
		if err := c.Bind(p); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updated, err := h.productsService.Update(ctx, p, id)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, updated)
	}
}

// GetByID godoc
// @Summary Get by id products
// @Description Get by id products handler
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id} [get]
func (h productsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "productsHandlers.GetByID")
		defer span.Finish()

		id := c.Param("id")

		res, err := h.productsService.GetByID(ctx, id)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetByID godoc
// @Summary Get all products
// @Description Get all products handler
// @Tags Products
// @Accept json
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.ProductList
// @Router /products [get]
func (h productsHandlers) All() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "productsHandlers.GetByID")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		itemsByID, err := h.productsService.All(ctx, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, itemsByID)
	}
}

// Delete godoc
// @Summary Delete products
// @Description Delete by id products handler
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string	"ok"
// @Router /products/{id} [delete]
func (h productsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "productsHandlers.Delete")
		defer span.Finish()

		id := c.Param("id")

		if err := h.productsService.Delete(ctx, id); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

package http

import (
	"net/http"

	"github.com/DLzer/go-echo-boilerplate/config"
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/orders"
	"github.com/DLzer/go-echo-boilerplate/pkg/httpErrors"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

// Orders Handler
type ordersHandlers struct {
	cfg           *config.Config
	ordersService orders.Service
	logger        logger.Logger
}

// NewOrdersHandlers Orders handlers constructor
func NewOrdersHandlers(cfg *config.Config, ordersService orders.Service, logger logger.Logger) orders.Handlers {
	return &ordersHandlers{cfg: cfg, ordersService: ordersService, logger: logger}
}

// Create godoc
// @Summary Create orders
// @Description Create orders handler
// @Tags Orders
// @Accept json
// @Produce json
// @Success 201 {object} models.OrderResponse
// @Router /orders/create [post]
func (h ordersHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "ordersHandlers.Create")
		defer span.Finish()

		p := &models.OrderRequest{}
		if err := c.Bind(p); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		created, err := h.ordersService.Create(ctx, p)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, created)
	}
}

// Create godoc
// @Summary Update orders
// @Description Update orders handler
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 201 {object} models.OrderResponse
// @Router /orders/update/{id} [patch]
func (h ordersHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "ordersHandlers.Update")
		defer span.Finish()

		id := c.Param("id")

		p := &models.OrderRequest{}
		if err := c.Bind(p); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updated, err := h.ordersService.Update(ctx, p, id)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, updated)
	}
}

// GetByID godoc
// @Summary Get by id orders
// @Description Get by id orders handler
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.OrderResponse
// @Router /orders/{id} [get]
func (h ordersHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "ordersHandlers.GetByID")
		defer span.Finish()

		id := c.Param("id")

		res, err := h.ordersService.GetByID(ctx, id)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetByID godoc
// @Summary Get all orders
// @Description Get all orders handler
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.OrderList
// @Router /orders [get]
func (h ordersHandlers) All() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "ordersHandlers.GetByID")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		itemsByID, err := h.ordersService.All(ctx, pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, itemsByID)
	}
}

// Delete godoc
// @Summary Delete orders
// @Description Delete by id orders handler
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string	"ok"
// @Router /orders/{id} [delete]
func (h ordersHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "ordersHandlers.Delete")
		defer span.Finish()

		id := c.Param("id")

		if err := h.ordersService.Delete(ctx, id); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

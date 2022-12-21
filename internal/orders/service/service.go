package service

import (
	"context"

	"github.com/DLzer/go-echo-boilerplate/config"
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/orders"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

// Orders Service
type ordersService struct {
	cfg        *config.Config
	ordersRepo orders.Repository
	logger     logger.Logger
}

// Orders Service constructor
func NewOrdersService(cfg *config.Config, ordersRepo orders.Repository, logger logger.Logger) orders.Service {
	return &ordersService{cfg: cfg, ordersRepo: ordersRepo, logger: logger}
}

// Create order
func (u *ordersService) Create(ctx context.Context, order *models.OrderRequest) (*models.OrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersService.Create")
	defer span.Finish()

	p, err := u.ordersRepo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return p, err
}

// Update order
func (u *ordersService) Update(ctx context.Context, order *models.OrderRequest, id string) (*models.OrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersService.Update")
	defer span.Finish()

	updated, err := u.ordersRepo.Update(ctx, order, id)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// Get order by  id
func (u *ordersService) GetByID(ctx context.Context, id string) (*models.OrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersService.GetByID")
	defer span.Finish()

	p, err := u.ordersRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Get all orders
func (u *ordersService) All(ctx context.Context, pq *utils.PaginationQuery) (*models.OrderList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersService.GetAll")
	defer span.Finish()

	p, err := u.ordersRepo.All(ctx, pq)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Delete order
func (u *ordersService) Delete(ctx context.Context, id string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersService.Delete")
	defer span.Finish()

	if err := u.ordersRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

package orders

import (
	"context"

	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
)

type Service interface {
	Create(ctx context.Context, order *models.OrderRequest) (*models.OrderResponse, error)
	Update(ctx context.Context, order *models.OrderRequest, id string) (*models.OrderResponse, error)
	GetByID(ctx context.Context, id string) (*models.OrderResponse, error)
	All(ctx context.Context, pq *utils.PaginationQuery) (*models.OrderList, error)
	Delete(ctx context.Context, id string) error
}

package products

import (
	"context"

	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
)

type Repository interface {
	Create(ctx context.Context, product *models.ProductRequest) (*models.ProductResponse, error)
	Update(ctx context.Context, product *models.ProductRequest, id string) (*models.ProductResponse, error)
	GetByID(ctx context.Context, id string) (*models.ProductResponse, error)
	All(ctx context.Context, pq *utils.PaginationQuery) (*models.ProductList, error)
	Delete(ctx context.Context, id string) error
}

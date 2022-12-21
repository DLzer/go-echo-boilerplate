package service

import (
	"context"

	"github.com/DLzer/go-echo-boilerplate/config"
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/products"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

// Products Service
type productsService struct {
	cfg          *config.Config
	productsRepo products.Repository
	logger       logger.Logger
}

// Products Service constructor
func NewProductsService(cfg *config.Config, productsRepo products.Repository, logger logger.Logger) products.Service {
	return &productsService{cfg: cfg, productsRepo: productsRepo, logger: logger}
}

// Create products
func (u *productsService) Create(ctx context.Context, product *models.ProductRequest) (*models.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsService.Create")
	defer span.Finish()

	p, err := u.productsRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return p, err
}

// Update products
func (u *productsService) Update(ctx context.Context, product *models.ProductRequest, id string) (*models.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsService.Update")
	defer span.Finish()

	updated, err := u.productsRepo.Update(ctx, product, id)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// Get products by  id
func (u *productsService) GetByID(ctx context.Context, id string) (*models.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsService.GetByID")
	defer span.Finish()

	p, err := u.productsRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Get all products
func (u *productsService) All(ctx context.Context, pq *utils.PaginationQuery) (*models.ProductList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsService.GetAll")
	defer span.Finish()

	p, err := u.productsRepo.All(ctx, pq)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Delete products
func (u *productsService) Delete(ctx context.Context, id string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsService.Delete")
	defer span.Finish()

	if err := u.productsRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

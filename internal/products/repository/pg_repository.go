package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/products"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type productsRepo struct {
	db *sqlx.DB
}

// Products repository constructor
func NewProductsRepository(db *sqlx.DB) products.Repository {
	return &productsRepo{db: db}
}

// Create products
func (r *productsRepo) Create(ctx context.Context, product *models.ProductRequest) (*models.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsRepo.Create")
	defer span.Finish()

	dt := time.Now()

	var p models.ProductResponse
	if err := r.db.QueryRowxContext(
		ctx,
		createProduct,
		&product.Name,
		&product.Description,
		&dt,
		&dt,
		0,
	).StructScan(&p); err != nil {
		return nil, errors.Wrap(err, "productsRepo.Create.QueryRowContext")
	}

	return &p, nil
}

// Update products item
func (r *productsRepo) Update(ctx context.Context, product *models.ProductRequest, id string) (*models.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsRepo.Update")
	defer span.Finish()

	var p models.ProductResponse

	dt := time.Now()
	if err := r.db.QueryRowxContext(
		ctx,
		updateProducts,
		&product.Name,
		&product.Description,
		&dt,
		id,
	).StructScan(&p); err != nil {
		return nil, errors.Wrap(err, "productsRepo.Update.QueryRowxContext")
	}

	return &p, nil
}

// Get single product by id
func (r *productsRepo) GetByID(ctx context.Context, id string) (*models.ProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsRepo.GetByID")
	defer span.Finish()

	var p models.ProductResponse
	if err := r.db.GetContext(ctx, &p, getProductById, id); err != nil {
		return nil, errors.Wrap(err, "productsRepo.GetByID.GetContext")
	}

	return &p, nil
}

// Get all products
func (r *productsRepo) All(ctx context.Context, pq *utils.PaginationQuery) (*models.ProductList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsRepo.GetAll")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, countProducts); err != nil {
		return nil, errors.Wrap(err, "productsRepo.GetAll.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.ProductList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Products:   make([]*models.ProductResponse, 0),
		}, nil
	}

	var list = make([]*models.ProductResponse, 0, pq.GetSize())
	rows, err := r.db.QueryxContext(ctx, allProducts, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "productsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.ProductResponse{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "productsRepo.GetAll.StructScan")
		}
		list = append(list, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "productsRepo.GetAll.rows.Err")
	}

	productList := &models.ProductList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Products:   list,
	}

	return productList, nil
}

// Delete products by id
func (r *productsRepo) Delete(ctx context.Context, id string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deleteProducts, id)
	if err != nil {
		return errors.Wrap(err, "productsRepo.Delete.ExecContent")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "productsRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "productsRepo.Delete.rowsAffected")
	}

	return nil
}

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/orders"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type ordersRepo struct {
	db *sqlx.DB
}

// Orders repository constructor
func NewOrdersRepository(db *sqlx.DB) orders.Repository {
	return &ordersRepo{db: db}
}

// Create orders
func (r *ordersRepo) Create(ctx context.Context, order *models.OrderRequest) (*models.OrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersRepo.Create")
	defer span.Finish()

	dt := time.Now()

	var o models.OrderResponse
	if err := r.db.QueryRowxContext(
		ctx,
		createOrder,
		&order.Total,
		&dt,
		&dt,
		0,
	).StructScan(&o); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.Create.QueryRowContext")
	}

	for _, x := range order.Products {
		if _, err := r.db.ExecContext(
			ctx,
			createOrderProductJoin,
			&o.ID,
			x,
		); err != nil {
			return nil, errors.Wrap(err, "ordersRepo.Create.QueryRowContext")
		}
	}

	return &o, nil
}

// Update orders item
func (r *ordersRepo) Update(ctx context.Context, order *models.OrderRequest, id string) (*models.OrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersRepo.Update")
	defer span.Finish()

	var o models.OrderResponse

	dt := time.Now()
	if err := r.db.QueryRowxContext(
		ctx,
		updateOrders,
		&order.Total,
		&dt,
		id,
	).StructScan(&o); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.Update.QueryRowxContext")
	}

	return &o, nil
}

// Get single orders by id
func (r *ordersRepo) GetByID(ctx context.Context, id string) (*models.OrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersRepo.GetByID")
	defer span.Finish()

	var o models.OrderResponse
	if err := r.db.GetContext(ctx, &o, getOrderById, id); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.GetByID.GetContext")
	}

	return &o, nil
}

// Get all orders
func (r *ordersRepo) All(ctx context.Context, pq *utils.PaginationQuery) (*models.OrderList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersRepo.GetAll")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, countOrders); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.GetAll.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.OrderList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Orders:     make([]*models.OrderResponse, 0),
		}, nil
	}

	var list = make([]*models.OrderResponse, 0, pq.GetSize())
	rows, err := r.db.QueryxContext(ctx, allOrders, pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "ordersRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	for rows.Next() {
		n := &models.OrderResponse{}
		if err = rows.StructScan(n); err != nil {
			return nil, errors.Wrap(err, "ordersRepo.GetAll.StructScan")
		}
		list = append(list, n)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.GetAll.rows.Err")
	}

	orderList := &models.OrderList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Orders:     list,
	}

	return orderList, nil
}

// Delete orders by id
func (r *ordersRepo) Delete(ctx context.Context, id string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ordersRepo.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deleteOrders, id)
	if err != nil {
		return errors.Wrap(err, "ordersRepo.Delete.ExecContent")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "ordersRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "ordersRepo.Delete.rowsAffected")
	}

	return nil
}

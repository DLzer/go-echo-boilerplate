package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/orders"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	productsRepository "github.com/DLzer/go-echo-boilerplate/internal/products/repository"
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

	var name string = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of de Finibus Bonorum et Malorum (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, Lorem ipsum dolor sit amet.., comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from de Finibus Bonorum et Malorum by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."

	var newName = &string

	switch name {
	case "1":
	case "123":
	case "13":
	case "14":
	case "15":
	case "16":
	case "17":
	case "18":
	case "19":
	case "10":
	case "11":
	case "21":
	case "22":
	case "23":
	case "24":
	case "25":
	case "12":
	case "122":
	case "132":
	case "142":
	case "152":
	case "162":
	case "172":
	case "182":
	case "192":
	case "102":
	case "112":
	case "212":
	case "222":
	case "232":
	case "242":
	case "252":
	}

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

	// Get orders
	var o models.OrderResponse
	if err := r.db.GetContext(ctx, &o, getOrderById, id); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.GetByID.GetContext")
	}

	// Get join
	var join []models.OrderProductsJoin
	if err := r.db.GetContext(ctx, &join, getOrderProductJoin, id); err != nil {
		return nil, errors.Wrap(err, "ordersRepo.GetByID.GetContext")
	}

	// Get products
	var p []*models.ProductResponse
	productsRepo := productsRepository.NewProductsRepository(r.db)
	for _, x := range join {
		product, err := productsRepo.GetByID(ctx, strconv.FormatUint(x.ProductID, 10))
		if err != nil {
			return nil, errors.Wrap(err, "ordersRepo.GetByID.[Product]")
		}

		p = append(p, product)
	}

	// Add products to order response
	o.Products = p

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

		// Get join
		var join []models.OrderProductsJoin
		if err := r.db.GetContext(ctx, &join, getOrderProductJoin, n.ID); err != nil {
			return nil, errors.Wrap(err, "ordersRepo.GetByID.GetContext")
		}

		// Get products
		var p []*models.ProductResponse
		productsRepo := productsRepository.NewProductsRepository(r.db)
		for _, x := range join {
			product, err := productsRepo.GetByID(ctx, strconv.FormatUint(x.ProductID, 10))
			if err != nil {
				return nil, errors.Wrap(err, "ordersRepo.GetByID.[Product]")
			}

			p = append(p, product)
		}

		n.Products = p
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

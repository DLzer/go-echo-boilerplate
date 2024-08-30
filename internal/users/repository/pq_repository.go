package repository

import (
	"context"
	"fmt"

	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/users"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository
type users_pg_repo struct {
	db *pgxpool.Pool
}

// Repository Constructor
func NewUsersPgRepo(db *pgxpool.Pool) users.Repository {
	return &users_pg_repo{db: db}
}

// Creates a user
func (r *users_pg_repo) Create(user *models.UserCreate) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), createQuery,
		user.UUID,
		user.Email,
		user.FirstName,
		user.LastName,
	)
	if err != nil {
		return nil, fmt.Errorf("users.PgRepository.Create: %v", err)
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.User])
	if err != nil {
		return nil, fmt.Errorf("users.Repository.Create.CollectOneRow: %v", err)
	}

	return &t, nil
}

// Update a user
func (r *users_pg_repo) Update(user *models.UserUpdate, user_uuid uuid.UUID) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), updateQuery,
		user.Email,
		user.FirstName,
		user.LastName,
		user_uuid,
	)
	if err != nil {
		return nil, fmt.Errorf("users.PgRepository.Update: %v", err)
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.User])
	if err != nil {
		return nil, fmt.Errorf("users.Repository.Update.CollectOneRow: %v", err)
	}

	return &t, nil
}

// Delete a user
func (r *users_pg_repo) Delete(user_id uuid.UUID) (bool, error) {
	_, err := r.db.Exec(context.Background(), deleteQuery, user_id)
	if err != nil {
		return false, fmt.Errorf("users.PgRepository.Delete: %v", err)
	}

	return true, nil
}

// Accepts a UUID, returns a user
func (r *users_pg_repo) GetByID(uuid uuid.UUID) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), getByIDQuery, uuid)
	if err != nil {
		return nil, fmt.Errorf("users.Repository.GetByID: %v", err)
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.User])
	if err != nil {
		return nil, fmt.Errorf("users.Repository.GetByID.GetByID: %v", err)
	}

	return &t, nil
}

// Accepts a PaginationQuery as parameter and returns a list of users, and an error if there is one
func (r *users_pg_repo) GetList(pq *utils.PaginationQuery) (*models.PaginationListType[models.User], error) {
	// Get a total count for reference
	var total_count int
	if err := r.db.QueryRow(context.Background(), countQuery).Scan(&total_count); err != nil {
		return nil, fmt.Errorf("users.PgRepository.GetList.Count: %v", err)
	}

	// If total count == 0 respond with a generic empty list
	if total_count == 0 {
		return &models.PaginationListType[models.User]{
			TotalCount: total_count,
			TotalPages: utils.GetTotalPages(total_count, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), total_count, pq.GetSize()),
			Values:     make([]models.User, 0),
		}, nil
	}

	// Make a []slice with a size determined by pq.GetSize
	var err error
	rows, err := r.db.Query(context.Background(), listQuery, pq.GetOrderBy(), pq.GetOffset(), pq.GetLimit())
	if err != nil {
		return nil, fmt.Errorf("users.PgRepository.GetList.Query: %v", err)
	}
	defer rows.Close()

	// Scan each row into a temp struct/memory slot --> Then append it to the final list
	var users []models.User
	if users, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.User]); err != nil {
		return nil, fmt.Errorf("users.PgRepository.GetList.CollectRows: %v", err)
	}

	// Create + Return our final model, and a NIL error
	return &models.PaginationListType[models.User]{
		TotalCount: total_count,
		TotalPages: utils.GetTotalPages(total_count, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), total_count, pq.GetSize()),
		Values:     users,
	}, nil
}

// Accepts a string to perform a vector search, returns a business list model or an error
func (r *users_pg_repo) Search(searchString string) (*models.PaginationListType[models.User], error) {
	searchFormat := fmt.Sprintf("%s:*", searchString)

	// Make a []slice with a size determined by pq.GetSize
	rows, err := r.db.Query(context.Background(), nameSearchQuery, searchFormat)
	if err != nil {
		return nil, fmt.Errorf("users.Repository.Search.Query: %v", err)
	}
	defer rows.Close()

	// Scan each row into a temp struct/memory slot --> Then append it to the final list
	var users []models.User
	if users, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.User]); err != nil {
		return nil, fmt.Errorf("users.PgRepository.Search.CollectRows: %v", err)
	}

	// Create + Return our final model, and a NIL error
	return &models.PaginationListType[models.User]{
		TotalCount: 10,
		TotalPages: 1,
		Page:       1,
		Size:       10,
		HasMore:    false,
		Values:     users,
	}, nil
}

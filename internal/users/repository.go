package users

import (
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/google/uuid"
)

type Repository interface {
	// Create
	Create(user *models.UserCreate) (*models.User, error)
	// Update
	Update(user *models.UserUpdate, user_uuid uuid.UUID) (*models.User, error)
	// Delete
	Delete(user_id uuid.UUID) (bool, error)
	// Get By ID
	GetByID(uuid uuid.UUID) (*models.User, error)
	// Get List
	GetList(pq *utils.PaginationQuery) (*models.PaginationListType[models.User], error)
	// Name Search
	Search(searchString string) (*models.PaginationListType[models.User], error)
}

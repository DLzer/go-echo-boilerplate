package service

import (
	"fmt"

	"github.com/DLzer/go-echo-boilerplate/internal/config"
	"github.com/DLzer/go-echo-boilerplate/internal/models"
	"github.com/DLzer/go-echo-boilerplate/internal/users"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/google/uuid"
)

// User Service
type user_service struct {
	cfg          *config.Config
	user_pg_repo users.Repository
	logger       logger.Logger
}

// User Constructor
func NewUsersService(cfg *config.Config, user_pg_repo users.Repository, logger logger.Logger) users.Service {
	return &user_service{
		cfg:          cfg,
		user_pg_repo: user_pg_repo,
		logger:       logger,
	}
}

// Creates a user
func (s *user_service) Create(user *models.UserCreate) (*models.User, error) {
	created_user, err := s.user_pg_repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("users.Service.Create: %v", err)
	}

	return created_user, nil
}

// Update a user
func (s *user_service) Update(user *models.UserUpdate, user_uuid uuid.UUID) (*models.User, error) {
	t, err := s.user_pg_repo.Update(user, user_uuid)
	if err != nil {
		return nil, fmt.Errorf("users.Service.Update: %v", err)
	}

	return t, nil
}

// Delete a user
func (s *user_service) Delete(user_id uuid.UUID) (bool, error) {
	d, err := s.user_pg_repo.Delete(user_id)
	if err != nil {
		return false, fmt.Errorf("users.Service.Delete: %v", err)
	}

	return d, nil
}

// Get User by ID
func (s *user_service) GetByID(user_id uuid.UUID) (*models.User, error) {
	t, err := s.user_pg_repo.GetByID(user_id)
	if err != nil {
		return nil, fmt.Errorf("users.Service.GetByID: %v", err)
	}

	return t, nil
}

// Accepts a PaginationQuery and returns a list of users
func (s *user_service) GetList(pq *utils.PaginationQuery) (*models.PaginationListType[models.User], error) {
	t, err := s.user_pg_repo.GetList(pq)
	if err != nil {
		return nil, fmt.Errorf("users.Service.GetList: %v", err)
	}

	return t, nil
}

// Accepts a string to perform a vector search, returns a user list model or an error
func (s *user_service) Search(searchString string) (*models.PaginationListType[models.User], error) {
	searchResult, err := s.user_pg_repo.Search(searchString)
	if err != nil {
		return nil, fmt.Errorf("users.Service.Search: %v", err)
	}

	return searchResult, nil
}

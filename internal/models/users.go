package models

import (
	"time"

	"github.com/google/uuid"
)

// User with Roles
// @Description User base model with Roles array
type User struct {
	UUID      uuid.UUID  `db:"uuid" json:"uuid,omitempty" example:"adcce0b7-0b38-4bd3-bfa1-d9bf7c4c79b4"`
	Email     *string    `db:"email" json:"email,omitempty" example:"someuser@name.com"`
	FirstName *string    `db:"first_name" json:"first_name,omitempty" example:"Geoff"`
	LastName  *string    `db:"last_name" json:"last_name,omitempty" example:"Goldblum"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty" example:"2024-01-01 00:01:22"`
	IsDeleted bool       `db:"is_deleted" json:"is_deleted,omitempty" example:"false"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty" example:"2024-01-01 00:01:22"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty" example:"2024-01-01 00:01:22"`
	Roles     []string   `db:"roles" json:"roles" example:"[USER]"`
}

// User Create
// @Description User Create Model
type UserCreate struct {
	UUID      string   `db:"uuid" json:"uuid,omitempty" example:"adcce0b7-0b38-4bd3-bfa1-d9bf7c4c79b4"`
	Email     *string  `db:"email" json:"email,omitempty" example:"someuser@name.com"`
	FirstName *string  `db:"first_name" json:"first_name,omitempty" example:"Geoff"`
	LastName  *string  `db:"last_name" json:"last_name,omitempty" example:"Goldblum"`
	Roles     []string `db:"roles" json:"roles" example:"[USER]"`
}

// User Update
// @Description User Update Model
type UserUpdate struct {
	UUID      string   `db:"uuid" json:"uuid,omitempty" example:"adcce0b7-0b38-4bd3-bfa1-d9bf7c4c79b4"`
	Email     *string  `db:"email" json:"email,omitempty" example:"someuser@name.com"`
	FirstName *string  `db:"first_name" json:"first_name,omitempty" example:"Geoff"`
	LastName  *string  `db:"last_name" json:"last_name,omitempty" example:"Goldblum"`
	Roles     []string `db:"roles" json:"roles" example:"[USER]"`
}

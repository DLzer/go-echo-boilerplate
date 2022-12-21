package models

import "time"

// ProductRequest is used to represent a product creation model
type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ProductResponse is used to represent a product response model
type ProductResponse struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Created     *time.Time `json:"created,omitempty"`
	Updated     *time.Time `json:"updated,omitempty"`
	Deleted     bool       `json:"deleted,omitempty"`
}

// ProductList is used to represent a pagination query for products
type ProductList struct {
	TotalCount int                `json:"total_count"`
	TotalPages int                `json:"total_pages"`
	Page       int                `json:"page"`
	Size       int                `json:"size"`
	HasMore    bool               `json:"has_more"`
	Products   []*ProductResponse `json:"products"`
}

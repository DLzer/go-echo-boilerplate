package models

import (
	"time"

	"github.com/lib/pq"
)

// OrderRequest is used to represent a order creation model
type OrderRequest struct {
	Total    float64        `json:"name"`
	Products pq.StringArray `json:"products"`
}

// OrderResponse is used to represent a order response model
type OrderResponse struct {
	ID       uint64             `json:"id"`
	Total    float64            `json:"name"`
	Products []*ProductResponse `json:"products,omitempty"`
	Created  *time.Time         `json:"created,omitempty"`
	Updated  *time.Time         `json:"updated,omitempty"`
	Deleted  bool               `json:"deleted,omitempty"`
}

// OrderProductsJoin is used to represent a join between orders and products
type OrderProductsJoin struct {
	ID        uint64
	OrderID   uint64
	ProductID uint64
}

// OrderList is used to represent a pagination query for orders
type OrderList struct {
	TotalCount int              `json:"total_count"`
	TotalPages int              `json:"total_pages"`
	Page       int              `json:"page"`
	Size       int              `json:"size"`
	HasMore    bool             `json:"has_more"`
	Orders     []*OrderResponse `json:"orders"`
}

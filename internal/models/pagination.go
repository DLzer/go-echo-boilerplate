package models

// Pagination List Type Response Model
// @Description Pagination List Type Response Model
type PaginationListType[T any] struct {
	TotalCount int  `json:"total_count"`
	TotalPages int  `json:"total_pages"`
	Page       int  `json:"page"`
	Size       int  `json:"size"`
	HasMore    bool `json:"has_more"`
	Values     []T  `json:"values"`
}

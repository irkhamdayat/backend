package response

import (
	"math"
)

// MetaResp represents metadata for pagination.
type MetaResp struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

// PaginationResp is a generic response structure for paginated data.
type PaginationResp[T any] struct {
	Meta  MetaResp `json:"meta"`
	Items []T      `json:"items"`
}

// GetTotalPages calculates the total number of pages for pagination.
func (meta *MetaResp) GetTotalPages(totalItems int64) int64 {
	if totalItems == 0 || meta.Limit == 0 {
		return 0
	}
	return int64(math.Ceil(float64(totalItems) / float64(meta.Limit)))
}

// NewPaginationResp creates a new PaginationResp instance with the given parameters.
func NewPaginationResp[T any](items []T, totalItems int64, meta MetaResp) *PaginationResp[T] {
	return &PaginationResp[T]{
		Meta: MetaResp{
			Page:       meta.Page,
			Limit:      meta.Limit,
			TotalItems: totalItems,
			TotalPages: meta.GetTotalPages(totalItems),
		},
		Items: items,
	}
}

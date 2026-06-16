package model

type WebResponse[T any] struct {
	Data   T                   `json:"data,omitempty"`
	Paging *PaginationResponse `json:"paging,omitempty"`
	Errors any                 `json:"errors,omitempty"`
}

type PaginationResponse struct {
	Page       int   `json:"page"`
	TotalPages int   `json:"total_pages"`
	TotalItems int64 `json:"total_items"`
}

type PaginationRequest struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
}

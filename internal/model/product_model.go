package model

import "time"

type ProductResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Stock     int64     `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=255"`
	Price int64  `json:"price" validate:"required,min=0"`
	Stock int64  `json:"stock" validate:"required,min=0"`
}

type UpdateProductRequest struct {
	ID    int64  `json:"-" validate:"required,min=1"`
	Name  string `json:"name" validate:"required,min=1,max=255"`
	Price int64  `json:"price" validate:"required,min=0"`
	Stock int64  `json:"stock" validate:"required,min=0"`
}

type PatchProductRequest struct {
	ID    int64   `json:"-" validate:"required,min=1"`
	Name  *string `json:"name" validate:"omitempty,min=1,max=255"`
	Price *int64  `json:"price" validate:"omitempty,min=0"`
	Stock *int64  `json:"stock" validate:"omitempty,min=0"`
}
package repository

import (
	"context"
	"toko-api-fiber/internal/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, entity *entity.Product) error
	Update(ctx context.Context, entity *entity.Product) error
	Delete(ctx context.Context, entity *entity.Product) error
	GetAll(ctx context.Context, page, size int) ([]*entity.Product, int64, int, error)
	GetByID(ctx context.Context, id int64) (*entity.Product, error)
	Patch(ctx context.Context, entity *entity.Product, fields map[string]any) error
}

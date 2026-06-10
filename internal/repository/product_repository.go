package repository

import (
	"context"
	"toko-api-fiber/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, entity *entity.Product) error
	Update(ctx context.Context, tx *gorm.DB, entity *entity.Product) error
	Delete(ctx context.Context, tx *gorm.DB, entity *entity.Product) error
	GetAll(ctx context.Context, tx *gorm.DB) ([]entity.Product, error)
	GetByID(ctx context.Context, tx *gorm.DB, id int64) (*entity.Product, error)
}

package repository

import (
	"context"
	"toko-api-fiber/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewProductRepository(db *gorm.DB, log *logrus.Logger) ProductRepository {
	return &ProductRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, entity *entity.Product) error {
	return tx.WithContext(ctx).Create(entity).Error
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, entity *entity.Product) error {
	return tx.WithContext(ctx).Updates(entity).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, entity *entity.Product) error {
	return tx.WithContext(ctx).Delete(entity).Error
}

func (r *ProductRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) ([]*entity.Product, error) {
	var products []*entity.Product
	return products, tx.WithContext(ctx).Find(&products).Error
}

func (r *ProductRepositoryImpl) GetByID(ctx context.Context, tx *gorm.DB, id int64) (*entity.Product, error) {
	var entity *entity.Product
	return entity, tx.WithContext(ctx).First(&entity, "id = ?", id).Error
}

func (r *ProductRepositoryImpl) Patch(ctx context.Context, tx *gorm.DB, entity *entity.Product, fields map[string]any) error {
	return tx.WithContext(ctx).Model(entity).Where("id = ?", entity.ID).Updates(fields).Scan(entity).Error
}
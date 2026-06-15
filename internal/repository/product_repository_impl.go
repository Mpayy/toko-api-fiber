package repository

import (
	"context"
	"errors"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/exception"

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

func (r *ProductRepositoryImpl) GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := GetTxFromContext(ctx); ok {
		return tx
	}
	return r.DB
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, entity *entity.Product) error {
	err := r.GetTx(ctx).Create(entity).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, entity *entity.Product) error {
	err := r.GetTx(ctx).Updates(entity).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, entity *entity.Product) error {
	err := r.GetTx(ctx).Delete(entity).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product
	err := r.GetTx(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryImpl) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
	var entity *entity.Product

	if err := r.GetTx(ctx).First(&entity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, err
	}

	return entity, nil
}

func (r *ProductRepositoryImpl) Patch(ctx context.Context, entity *entity.Product, fields map[string]any) error {
	err := r.GetTx(ctx).Model(entity).Where("id = ?", entity.ID).Updates(fields).Scan(entity).Error
	if err != nil {
		return err
	}

	return nil
}

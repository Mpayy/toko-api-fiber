package repository

import (
	"context"
	"errors"
	"time"
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

func (r *ProductRepositoryImpl) logIfSlowQuery(startTime time.Time, action string, queryInfo any) {
	latency := time.Since(startTime)
	threshold := time.Duration(200 * time.Millisecond)
	if latency > threshold {
		r.Log.WithFields(logrus.Fields{
			"layer":   "REPOSITORY",
			"action":  action,
			"target":  queryInfo,
			"latency": latency.String(),
		}).Warn("Database slow query detected!")
	}
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, entity *entity.Product) error {
	defer r.logIfSlowQuery(time.Now(), "Create Product", entity)
	return r.GetTx(ctx).Create(entity).Error
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, entity *entity.Product) error {
	defer r.logIfSlowQuery(time.Now(), "Update Product", entity)
	return r.GetTx(ctx).Updates(entity).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, entity *entity.Product) error {
	defer r.logIfSlowQuery(time.Now(), "Delete Product", entity)
	return r.GetTx(ctx).Delete(entity).Error
}

func (r *ProductRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Product, error) {
	defer r.logIfSlowQuery(time.Now(), "Get All Products", nil)
	var products []*entity.Product
	return products, r.GetTx(ctx).Find(&products).Error
}

func (r *ProductRepositoryImpl) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
	defer r.logIfSlowQuery(time.Now(), "Get Product By ID", id)
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
	defer r.logIfSlowQuery(time.Now(), "Patch Product", entity)
	return r.GetTx(ctx).Model(entity).Where("id = ?", entity.ID).Updates(fields).Scan(entity).Error
}

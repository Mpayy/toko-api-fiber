package repository

import (
	"context"
	"errors"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/exception"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) UserRepository {
	return &UserRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *UserRepositoryImpl) GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := GetTxFromContext(ctx); ok {
		return tx
	}
	return r.DB
}

func (r *UserRepositoryImpl) Create(ctx context.Context, entity *entity.User) error {
	err := r.GetTx(ctx).Create(entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exception.ErrDuplicatedEmail
		}
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, entity *entity.User) error {
	if err := r.GetTx(ctx).Model(entity).Select("*").Updates(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, entity *entity.User, id int64) error {
	if err := r.GetTx(ctx).Where("id = ?", id).First(entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, entity *entity.User, email string) error {
	if err := r.GetTx(ctx).Where("email = ?", email).First(entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByToken(ctx context.Context, entity *entity.User, token string) error {
	if err := r.GetTx(ctx).Where("token = ?", token).First(entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exception.ErrNotFound
		}
		return err
	}
	return nil
}

// func (r *UserRepositoryImpl) CountByEmail(ctx context.Context, email string) (int64, error) {
// 	var count int64
// 	if err := r.GetTx(ctx).Model(&entity.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

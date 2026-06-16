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
	if err := r.GetTx(ctx).Model(&entity).Select("*").Updates(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	var user *entity.User
	if err := r.GetTx(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	if err := r.GetTx(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByToken(ctx context.Context, token string) (*entity.User, error) {
	var user *entity.User
	if err := r.GetTx(ctx).Where("token = ?", token).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

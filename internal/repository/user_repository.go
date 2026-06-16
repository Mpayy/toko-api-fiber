package repository

import (
	"context"
	"toko-api-fiber/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, entity *entity.User) error
	Update(ctx context.Context, entity *entity.User) error
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByToken(ctx context.Context, token string) (*entity.User, error)
}
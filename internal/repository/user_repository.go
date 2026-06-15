package repository

import (
	"context"
	"toko-api-fiber/internal/entity"
)

type UserRepository interface {
	Create(ctx context.Context, entity *entity.User) error
	Update(ctx context.Context, entity *entity.User) error
	FindByID(ctx context.Context, entity *entity.User, id int64) error
	FindByEmail(ctx context.Context, entity *entity.User, email string) error
	FindByToken(ctx context.Context, entity *entity.User, token string) error
	// CountByEmail(ctx context.Context, email string) (int64, error)
}
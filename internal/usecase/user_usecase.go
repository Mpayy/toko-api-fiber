package usecase

import (
	"context"
	"toko-api-fiber/internal/model"
)

type UserUsecase interface {
	Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error)
	Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error)
	Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error)
	Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error)
	Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error)
}

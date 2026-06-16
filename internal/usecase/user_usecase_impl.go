package usecase

import (
	"context"
	"errors"
	"fmt"
	"toko-api-fiber/internal/entity"
	"toko-api-fiber/internal/exception"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/model/converter"
	"toko-api-fiber/internal/repository"
	"toko-api-fiber/internal/util"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	Transaction    repository.Transaction
	Log            *logrus.Logger
	UserRepository repository.UserRepository
	TokenUtil      util.TokenUtil
}

func NewUserUsecase(transaction repository.Transaction, log *logrus.Logger, userRepository repository.UserRepository, tokenUtil util.TokenUtil) UserUsecase {
	return &UserUsecaseImpl{
		Transaction:    transaction,
		Log:            log,
		UserRepository: userRepository,
		TokenUtil:      tokenUtil,
	}
}

func (u *UserUsecaseImpl) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	user, err := u.UserRepository.FindByToken(ctx, request.Token)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find user token: %w", err)
	}

	return &model.Auth{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserUsecaseImpl) Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &entity.User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Username: request.Username,
	}

	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		if errors.Is(err, exception.ErrDuplicatedEmail) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	return converter.ToUserResponse(user), nil
}

func (u *UserUsecaseImpl) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	user, err := u.UserRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrUnauthorized
		}
		return nil, fmt.Errorf("database error during login: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, exception.ErrUnauthorized
	}

	token, err := u.TokenUtil.CreateToken(&model.Auth{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}

	user.Token = token

	if err := u.UserRepository.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user token: %w", err)
	}

	return converter.ToUserTokenResponse(user), nil
}

func (u *UserUsecaseImpl) Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	user, err := u.UserRepository.FindByID(ctx, request.ID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return nil, exception.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return converter.ToUserResponse(user), nil
}

func (u *UserUsecaseImpl) Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error) {
	user, err := u.UserRepository.FindByID(ctx, request.ID)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return false, exception.ErrNotFound
		}
		return false, fmt.Errorf("failed to find user by id: %w", err)
	}

	user.Token = ""

	err = u.UserRepository.Update(ctx, user)
	if err != nil {
		return false, fmt.Errorf("failed to update user token: %w", err)
	}

	return true, nil
}

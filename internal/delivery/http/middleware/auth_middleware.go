package middleware

import (
	"strings"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/usecase"
	"toko-api-fiber/internal/util"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	Log         *logrus.Logger
	TokenUtil   util.TokenUtil
	UserUsecase usecase.UserUsecase
}

func NewAuthMiddleware(log *logrus.Logger, tokenUtil util.TokenUtil, userUsecase usecase.UserUsecase) *AuthMiddleware {
	return &AuthMiddleware{
		Log:         log,
		TokenUtil:   tokenUtil,
		UserUsecase: userUsecase,
	}
}

func (m *AuthMiddleware) Handle(ctx fiber.Ctx) error {
	authHeader := ctx.Get("Authorization", "NOT_FOUND")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request := &model.VerifyUserRequest{Token: tokenString}

	tokenRequest, err := m.TokenUtil.ParseToken(request.Token)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	tokenDb, err := m.UserUsecase.Verify(ctx.Context(), request)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if tokenDb.ID != tokenRequest.ID || tokenDb.Email != tokenRequest.Email || tokenDb.Username != tokenRequest.Username {
		return fiber.ErrUnauthorized
	}

	ctx.Locals("auth", tokenRequest)

	err = ctx.Next()

	m.Log.WithFields(logrus.Fields{
		"method": ctx.Method(),
		"path":   ctx.Path(),
		"status": ctx.Response().StatusCode(),
	}).Info("HTTP Request Completed")

	return err
}

func GetUser(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}

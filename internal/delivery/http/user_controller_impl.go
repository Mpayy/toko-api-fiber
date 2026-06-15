package http

import (
	"errors"
	"toko-api-fiber/internal/delivery/http/middleware"
	"toko-api-fiber/internal/exception"
	"toko-api-fiber/internal/model"
	"toko-api-fiber/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type UserControllerImpl struct {
	Log         *logrus.Logger
	UserUsecase usecase.UserUsecase
	Validator   *validator.Validate
}

func NewUserController(log *logrus.Logger, userUsecase usecase.UserUsecase, validator *validator.Validate) UserController {
	return &UserControllerImpl{
		Log:         log,
		UserUsecase: userUsecase,
		Validator:   validator,
	}
}

func (c *UserControllerImpl) Register(ctx fiber.Ctx) error {
	request := new(model.RegisterUserRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse request")
		return fiber.ErrBadRequest
	}

	err = c.Validator.Struct(request)
	if err != nil {
		fieldErrors := exception.ExtractValidationErrors(err)
		c.Log.WithFields(logrus.Fields{
			"errors": fieldErrors,
		}).Warn("Validation failed for register")
		return &exception.ValidationErrorWithFields{
			Errors: fieldErrors,
		}
	}

	response, err := c.UserUsecase.Register(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, exception.ErrDuplicatedEmail) {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserControllerImpl) Login(ctx fiber.Ctx) error {
	request := new(model.LoginUserRequest)

	err := ctx.Bind().JSON(request)
	if err != nil {
		c.Log.WithError(err).Warn("Failed to parse request")
		return fiber.ErrBadRequest
	}

	err = c.Validator.Struct(request)
	if err != nil {
		fieldErrors := exception.ExtractValidationErrors(err)
		c.Log.WithFields(logrus.Fields{
			"errors": fieldErrors,
		}).Warn("Validation failed for login")
		return &exception.ValidationErrorWithFields{
			Errors: fieldErrors,
		}
	}

	response, err := c.UserUsecase.Login(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, exception.ErrUnauthorized) {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserControllerImpl) Current(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.UserUsecase.Current(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserControllerImpl) Logout(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.UserUsecase.Logout(ctx.Context(), request)
	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}

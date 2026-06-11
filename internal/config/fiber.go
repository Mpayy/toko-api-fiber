package config

import (
	"errors"
	"toko-api-fiber/internal/model"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper, log *logrus.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: NewErrorHandler(log),
	})
	return app
}

func NewErrorHandler(log *logrus.Logger) fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		// code := fiber.StatusInternalServerError

		// if e, ok := err.(*fiber.Error); ok {
		// 	code = e.Code
		// }

		// return ctx.Status(code).JSON(fiber.Map{
		// 	"errors": err.Error(),
		// })

		log.WithError(err).Error("An error occurred")

		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		var fiberError *fiber.Error
		if errors.As(err, &fiberError) {
			code = fiberError.Code
			message = fiberError.Message
		} else if errors.Is(err, model.ErrNotFound) {
			code = fiber.StatusNotFound
			message = "Not Found"
		} else if errors.Is(err, model.ErrUnauthorized) {
			code = fiber.StatusUnauthorized
			message = "Unauthorized"
		} else if errors.Is(err, model.ErrValidation) {
			code = fiber.StatusBadRequest
			message = "Bad Request"
		}

		return ctx.Status(code).JSON(model.WebResponse[any]{
			Code:   code,
			Status: message,
			Errors: err.Error(),
		})
	}
}

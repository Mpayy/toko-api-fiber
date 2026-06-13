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
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"
		var errorMessage any = "INTERNAL SERVER ERROR"

		var clientError model.ClientError
		var fiberError *fiber.Error

		if errors.As(err, &clientError) {
			code = clientError.Code()
			message = clientError.Error()
			errorMessage = clientError.GetError()
		} else if errors.As(err, &fiberError) {
			code = fiberError.Code
			message = fiberError.Message
			errorMessage = fiberError.Error()
		}

		if code >= 500 {
			log.Error(err)
		}

		return ctx.Status(code).JSON(model.WebResponse[any]{
			Code:   code,
			Status: message,
			Errors: errorMessage,
		})
	}
}

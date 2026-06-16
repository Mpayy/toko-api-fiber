package middleware

import (
	"errors"
	"time"
	"toko-api-fiber/internal/exception"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func NewLoggingMiddleware(log *logrus.Logger) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()

		// Derive status code dari error karena ErrorHandler belum jalan
		status := ctx.Response().StatusCode()
		if err != nil {
			status = resolveStatusCode(err)
		}

		fields := logrus.Fields{
			"method":  ctx.Method(),
			"path":    ctx.Path(),
			"status":  status,
			"latency": time.Since(start).String(),
			"ip":      ctx.IP(),
		}

		if status >= 500 {
			log.WithFields(fields).Error("HTTP Request")
		} else if status >= 400 {
			log.WithFields(fields).Warn("HTTP Request")
		} else {
			log.WithFields(fields).Info("HTTP Request")
		}

		return err
	}
}

func resolveStatusCode(err error) int {
	// Cek custom ClientError (ValidationErrorWithFields, dll)
	var clientError exception.ClientError
	if errors.As(err, &clientError) {
		return clientError.Code()
	}

	// Cek fiber.Error (fiber.ErrBadRequest, fiber.ErrUnauthorized, dll)
	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		return fiberError.Code
	}

	// Default: Internal Server Error
	return fiber.StatusInternalServerError
}

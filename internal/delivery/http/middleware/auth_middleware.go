package middleware

import (
	"toko-api-fiber/internal/model"

	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthMiddleware struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewAuthMiddleware(log *logrus.Logger, config *viper.Viper) *AuthMiddleware {
	return &AuthMiddleware{
		Log:    log,
		Config: config,
	}
}

func (m *AuthMiddleware) Handle(ctx fiber.Ctx) error {
	method := ctx.Method()
	path := ctx.Path()
	ip := ctx.IP()

	m.Log.WithFields(logrus.Fields{
		"ip":     ip,
		"method": method,
		"path":   path,
	}).Info("Request Received")

	expectedKey := m.Config.GetString("APP_API_KEY")
	apiKey := ctx.Get("X-Api-Key")
	if apiKey != expectedKey {
		m.Log.WithFields(logrus.Fields{
			"method": method,
			"path":   path,
			"key":    apiKey,
		}).Error("Unauthorized")

		return &model.UnauthorizedError{
			Message: fiber.ErrUnauthorized.Message,
		}
	}

	err := ctx.Next()

	return err
}

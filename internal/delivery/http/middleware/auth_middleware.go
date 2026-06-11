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
	m.Log.WithFields(logrus.Fields{
		"method": ctx.Method(),
		"path":   ctx.Path(),
	}).Info("Request Received")

	expectedKey := m.Config.GetString("APP_API_KEY")
	apiKey := ctx.Get("X-Api-Key")
	if apiKey == expectedKey {
		return ctx.Next()
	}

	m.Log.WithFields(logrus.Fields{
		"method": ctx.Method(),
		"path":   ctx.Path(),
		"key":    apiKey,
	}).Error("Unauthorized")

	return model.ErrUnauthorized
}

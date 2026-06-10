package config

import (
	"toko-api-fiber/internal/delivery/http"
	"toko-api-fiber/internal/delivery/http/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	Fiber             *fiber.App
	ProductController http.ProductController
	AuthMiddleware    *middleware.AuthMiddleware
	Logger            *logrus.Logger
	Validation        *validator.Validate
	Config            *viper.Viper
	DB                *gorm.DB
}

func NewApp(
	app               *fiber.App,
	productController http.ProductController,
	authMiddleware    *middleware.AuthMiddleware,
	logger            *logrus.Logger,
	validation        *validator.Validate,
	config            *viper.Viper,
	db                *gorm.DB,
) *App {
	return &App{
		Fiber:             app,
		ProductController: productController,
		AuthMiddleware:    authMiddleware,
		Logger:            logger,
		Validation:        validation,
		Config:            config,
		DB:                db,
	}
}

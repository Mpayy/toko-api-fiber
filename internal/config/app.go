package config

import (
	"toko-api-fiber/internal/delivery/http"
	"toko-api-fiber/internal/delivery/http/middleware"
	"toko-api-fiber/internal/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type App struct {
	Fiber             *fiber.App
	ProductController http.ProductController
	UserController    http.UserController
	AuthMiddleware    *middleware.AuthMiddleware
	Logger            *logrus.Logger
	Validation        *validator.Validate
	Config            *viper.Viper
	DB                *gorm.DB
	TokenUtil         util.TokenUtil
}

func NewApp(
	app *fiber.App,
	productController http.ProductController,
	userController http.UserController,
	authMiddleware *middleware.AuthMiddleware,
	logger *logrus.Logger,
	validation *validator.Validate,
	config *viper.Viper,
	db *gorm.DB,
	tokenUtil util.TokenUtil,
) *App {
	return &App{
		Fiber:             app,
		ProductController: productController,
		UserController:    userController,
		AuthMiddleware:    authMiddleware,
		Logger:            logger,
		Validation:        validation,
		Config:            config,
		DB:                db,
		TokenUtil:         tokenUtil,
	}
}

//go:build wireinject
// +build wireinject

package main

import (
	"toko-api-fiber/internal/config"
	"toko-api-fiber/internal/delivery/http"
	"toko-api-fiber/internal/delivery/http/middleware"
	"toko-api-fiber/internal/repository"
	"toko-api-fiber/internal/usecase"

	"github.com/google/wire"
)

var productSet = wire.NewSet(
	repository.NewProductRepository,
	usecase.NewProductUseCase,
	http.NewProductController,
)

func InitializedApp() *config.App {
	wire.Build(
		config.NewViper,
		config.NewLogrus,
		config.NewGorm,
		config.NewValidator,
		config.NewFiber,
		productSet,
		middleware.NewAuthMiddleware,
		config.NewApp,
	)

	return nil
}

//go:build wireinject
// +build wireinject

package main

import (
	"toko-api-fiber/internal/config"
	"toko-api-fiber/internal/delivery/http"
	"toko-api-fiber/internal/delivery/http/middleware"
	"toko-api-fiber/internal/repository"
	"toko-api-fiber/internal/usecase"
	"toko-api-fiber/internal/util"

	"github.com/google/wire"
)

var productSet = wire.NewSet(
	repository.NewProductRepository,
	usecase.NewProductUseCase,
	http.NewProductController,
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	usecase.NewUserUsecase,
	http.NewUserController,
)

func InitializedApp() *config.App {
	wire.Build(
		config.NewViper,
		config.NewLogrus,
		config.NewGorm,
		config.NewValidator,
		config.NewFiber,
		repository.NewTransaction,
		productSet,
		userSet,
		middleware.NewAuthMiddleware,
		util.NewTokenUtil,
		config.NewApp,
	)

	return nil
}

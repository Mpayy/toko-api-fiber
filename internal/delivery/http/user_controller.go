package http

import (
	"github.com/gofiber/fiber/v3"
)

type UserController interface {
	Register(ctx fiber.Ctx) error
	Login(ctx fiber.Ctx) error
	Current(ctx fiber.Ctx) error
	Logout(ctx fiber.Ctx) error
}
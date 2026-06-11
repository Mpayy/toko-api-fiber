package http

import "github.com/gofiber/fiber/v3"

type ProductController interface {
	Create(ctx fiber.Ctx) error
	Update(ctx fiber.Ctx) error
	Delete(ctx fiber.Ctx) error
	GetAll(ctx fiber.Ctx) error
	GetByID(ctx fiber.Ctx) error
	Patch(ctx fiber.Ctx) error
}

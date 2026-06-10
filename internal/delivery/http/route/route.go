package route

import (
	"toko-api-fiber/internal/delivery/http"
	"toko-api-fiber/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v3"
)

func NewRouter(app *fiber.App, productController http.ProductController, authMiddleware *middleware.AuthMiddleware) {
	api := app.Group("/api")

	api.Use(authMiddleware.Handle)

	api.Post("/products", productController.Create)
	api.Put("/products/:id", productController.Update)
	api.Delete("/products/:id", productController.Delete)
	api.Get("/products", productController.GetAll)
	api.Get("/products/:id", productController.GetByID)
}

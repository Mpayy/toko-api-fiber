package route

import (
	"toko-api-fiber/internal/delivery/http"
	"toko-api-fiber/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v3"
)

func NewRouter(app *fiber.App, userController http.UserController, productController http.ProductController, authMiddleware *middleware.AuthMiddleware) {
    api := app.Group("/api")

    api.Post("/users/register", userController.Register)
    api.Post("/users/login", userController.Login)

    api.Get("/products", productController.GetAll)
    api.Get("/products/:id", productController.GetByID)

    protectedApi := api.Group("") 
    protectedApi.Use(authMiddleware.Handle)

    protectedApi.Get("/users/current", userController.Current)
    protectedApi.Delete("/users/logout", userController.Logout)

    protectedApi.Post("/products", productController.Create)
    protectedApi.Put("/products/:id", productController.Update)
    protectedApi.Patch("/products/:id", productController.Patch)
    protectedApi.Delete("/products/:id", productController.Delete)
}


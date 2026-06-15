package main

import (
	"fmt"
	"log"
	"toko-api-fiber/internal/delivery/http/route"
)

func main() {
	app := InitializedApp()

	route.NewRouter(
		app.Fiber,
		app.UserController,
		app.ProductController,
		app.AuthMiddleware,
	)

	host := app.Config.GetString("APP_HOST")
	port := app.Config.GetInt("APP_PORT")
	log.Printf("Server starting on: %s:%d\n", host, port)
	err := app.Fiber.Listen(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}

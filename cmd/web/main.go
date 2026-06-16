package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toko-api-fiber/internal/delivery/http/route"
)

func main() {
	app := InitializedApp()

	route.NewRouter(
		app.Fiber,
		app.UserController,
		app.ProductController,
		app.AuthMiddleware,
		app.Logger,
	)

	host := app.Config.GetString("APP_HOST")
	port := app.Config.GetInt("APP_PORT")
	log.Printf("Server starting on: %s:%d\n", host, port)

	go func() {
		if err := app.Fiber.Listen(fmt.Sprintf("%s:%d", host, port)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Fiber.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")

	db, err := app.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}

	log.Println("Database connection closed")
}

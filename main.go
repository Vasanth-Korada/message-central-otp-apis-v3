package main

import (
	"log"
	"message-central-v3/config"
	"message-central-v3/handler"
	"message-central-v3/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting the application...")

	// Load configuration
	config.LoadConfig()

	app := fiber.New()

	// Logger middleware
	app.Use(middleware.SetupLogger())

	// Public route (no auth middleware)
	app.Post("/auth/token", handler.GenerateToken)

	// Authenticated routes group
	api := app.Group("/app", middleware.AuthMiddleware())

	// Authenticated routes
	api.Post("/sendOtp", handler.SendOTP)
	api.Get("/verifyOtp", handler.ValidateOTP)

	// Start server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Server started successfully.")
}

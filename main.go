package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create instance of Fiber
	app := fiber.New()

	// Create HTTP handler
	app.Get("/testApi", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "go fiber first created api",
		})
	})

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

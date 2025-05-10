package main

import (
	"log"

	"anonyproof-protocol/backend/internal/adapters/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", handlers.HealthCheck)

	log.Fatal(app.Listen(":3000"))
}

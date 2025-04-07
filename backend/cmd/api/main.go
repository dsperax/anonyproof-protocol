package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/your/module/internal/adapters/handlers"
)

func main() {
    app := fiber.New()

    app.Get("/health", handlers.HealthCheck)

    log.Fatal(app.Listen(":3000"))
}

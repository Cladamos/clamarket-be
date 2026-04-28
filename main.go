package main

import (
	"log"

	"github.com/cladamos/clamarket-be/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/api/products", handlers.GetProducts)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

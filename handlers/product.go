package handlers

import (
	"github.com/cladamos/clamarket-be/repo"
	"github.com/gofiber/fiber/v3"
)

func GetProducts(r *repo.ProductRepository) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		products, err := r.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
		}
		return c.JSON(products)
	}
}

func GetProductByID(r *repo.ProductRepository) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		product, err := r.GetByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product"})
		}
		return c.JSON(product)
	}
}

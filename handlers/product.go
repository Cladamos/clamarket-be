package handlers

import (
	"github.com/cladamos/clamarket-be/repo"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetProducts(r *repo.ProductRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		products, err := r.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
		}
		return c.JSON(products)
	}
}

func GetProductByID(r *repo.ProductRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		product, err := r.GetByID(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product"})
		}
		return c.JSON(product)
	}
}

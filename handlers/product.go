package handlers

import (
	"github.com/cladamos/clamarket-be/models"
	"github.com/gofiber/fiber/v3"
)

var products = []models.Product{
	{
		ID:          "1",
		Name:        "Product 1",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Price:       100,
		Stock:       10,
		Image:       "image1.jpg",
		CategoryID:  "1",
		CreatedAt:   "2022-01-01",
		UpdatedAt:   "2022-01-01",
	},
	{
		ID:          "2",
		Name:        "Product 2",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Price:       200,
		Stock:       20,
		Image:       "image2.jpg",
		CategoryID:  "2",
		CreatedAt:   "2022-01-02",
		UpdatedAt:   "2022-01-02",
	},
	{
		ID:          "3",
		Name:        "Product 3",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Price:       300,
		Stock:       30,
		Image:       "image3.jpg",
		CategoryID:  "3",
		CreatedAt:   "2022-01-03",
		UpdatedAt:   "2022-01-03",
	},
}

func GetProducts(c fiber.Ctx) error {
	return c.JSON(products)
}

func GetProduct(c fiber.Ctx) error {
	id := c.Params("product_id")

	for _, product := range products {
		if product.ID == id {
			return c.JSON(product)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
}

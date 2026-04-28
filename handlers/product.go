package handlers

import (
	"github.com/cladamos/clamarket-be/models"
	"github.com/gofiber/fiber/v3"
)

var products = []models.Product{
	{
		ID:          1,
		Name:        "Product 1",
		Description: "Description 1",
		Price:       100,
		Stock:       10,
		Image:       "image1.jpg",
		CategoryID:  1,
		CreatedAt:   "2022-01-01",
		UpdatedAt:   "2022-01-01",
	},
	{
		ID:          2,
		Name:        "Product 2",
		Description: "Description 2",
		Price:       200,
		Stock:       20,
		Image:       "image2.jpg",
		CategoryID:  2,
		CreatedAt:   "2022-01-02",
		UpdatedAt:   "2022-01-02",
	},
	{
		ID:          3,
		Name:        "Product 3",
		Description: "Description 3",
		Price:       300,
		Stock:       30,
		Image:       "image3.jpg",
		CategoryID:  3,
		CreatedAt:   "2022-01-03",
		UpdatedAt:   "2022-01-03",
	},
}

func GetProducts(c fiber.Ctx) error {
	return c.JSON(products)
}

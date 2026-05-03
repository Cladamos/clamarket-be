package main

import (
	"log"
	"os"

	"github.com/cladamos/clamarket-be/handlers"
	"github.com/cladamos/clamarket-be/models"
	"github.com/cladamos/clamarket-be/repo"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	dsn := os.Getenv("DSN_STR")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.Product{}, &models.User{})

	productRepo := repo.NewProductRepository(db)
	userRepo := repo.NewUserRepository(db)

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Use(cors.New())
	app.Get("/api/products", handlers.GetProducts(productRepo))
	app.Get("/api/products/:id", handlers.GetProductByID(productRepo))
	app.Post("/api/users/register", handlers.Register(userRepo))
	app.Post("/api/users/login", handlers.Login(userRepo))

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

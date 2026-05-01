package handlers

import (
	"log"
	"strings"

	"github.com/cladamos/clamarket-be/models"
	"github.com/cladamos/clamarket-be/repo"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(r *repo.UserRepository) fiber.Handler {
	return func(c fiber.Ctx) error {

		u := new(models.RegisterRequest)
		if err := c.Bind().Body(u); err != nil {
			log.Println("Bind body failed: ", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Something went wrong"})
		}

		u.Name = strings.TrimSpace(u.Name)
		u.Email = strings.TrimSpace(u.Email)

		if validationErr := ValidateStruct(u); len(validationErr) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationErr})
		}

		isExist, err := r.IsExist(u.Email)
		if err != nil {
			log.Println("IsExist check failed: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
		}
		if isExist {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exist"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Hash password failed: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
		}
		if err = r.Create(&models.User{
			ID:       uuid.NewString(),
			Name:     u.Name,
			Email:    u.Email,
			Password: hashedPassword,
		}); err != nil {
			log.Println("Create user failed: ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
	}
}

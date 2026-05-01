package handlers

import (
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
			return NewBadRequestError("invalid request body")
		}

		u.Name = strings.TrimSpace(u.Name)
		u.Email = strings.TrimSpace(u.Email)

		if validationErr := ValidateStruct(u); len(validationErr) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error", "details": validationErr})
		}

		isExist, err := r.IsExist(u.Email)
		if err != nil {
			return NewInternalError("checking user existence", err)
		}
		if isExist {
			return NewConflictError("user already exists")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return NewInternalError("hashing password", err)
		}
		if err = r.Create(&models.User{
			ID:       uuid.NewString(),
			Name:     u.Name,
			Email:    u.Email,
			Password: hashedPassword,
		}); err != nil {
			return NewInternalError("creating user", err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
	}
}

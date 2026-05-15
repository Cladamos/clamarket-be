package handlers

import (
	"strings"

	"github.com/cladamos/clamarket-be/middleware"
	"github.com/cladamos/clamarket-be/models"
	"github.com/cladamos/clamarket-be/repo"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func Login(r *repo.UserRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		u := new(models.LoginRequest)
		if err := c.Bind().Body(u); err != nil {
			return NewBadRequestError("invalid request body")
		}

		u.Email = strings.TrimSpace(u.Email)

		if validationErr := ValidateStruct(u); len(validationErr) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "validation error", "details": validationErr})
		}

		user, err := r.GetByEmail(u.Email)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return NewUnauthorizedError("invalid credentials")
			}
			return NewInternalError("fetching user", err)
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(u.Password)); err != nil {
			return NewUnauthorizedError("invalid credentials")
		}
		token, err := middleware.GenerateToken(user.ID)
		if err != nil {
			return NewInternalError("generating token", err)
		}
		return c.JSON(fiber.Map{"message": "login success", "token": token})
	}
}

func GetMe(r *repo.UserRepository) fiber.Handler {
	return func(c fiber.Ctx) error {
		err := middleware.Auth(c)
		if err != nil {
			return err
		}

		userId := c.Locals("user_id").(string)
		user, err := r.GetByID(userId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return NewNotFoundError("user not found")
			}
			return NewInternalError("fetching user", err)
		}
		return c.JSON(fiber.Map{"message": "user found", "user": user})
	}
}

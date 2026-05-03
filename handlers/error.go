package handlers

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type AppError struct {
	StatusCode int
	Message    string
	Internal   error // the real error — logged, never sent to client
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return e.Internal.Error()
	}
	return e.Message
}

func NewBadRequestError(message string) *AppError {
	return &AppError{StatusCode: fiber.StatusBadRequest, Message: message}
}

func NewConflictError(message string) *AppError {
	return &AppError{StatusCode: fiber.StatusConflict, Message: message}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{StatusCode: fiber.StatusNotFound, Message: message}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{StatusCode: fiber.StatusUnauthorized, Message: message}
}

func NewInternalError(logMessage string, err error) *AppError {
	log.Printf("ERROR: %s: %v", logMessage, err)
	return &AppError{
		StatusCode: fiber.StatusInternalServerError,
		Message:    "something went wrong",
		Internal:   err,
	}
}

func ErrorHandler(c fiber.Ctx, err error) error {
	if appErr, ok := err.(*AppError); ok {
		return c.Status(appErr.StatusCode).JSON(fiber.Map{
			"error": appErr.Message,
		})
	}

	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"error": fiberErr.Message,
		})
	}

	log.Printf("ERROR: unexpected: %v", err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "something went wrong",
	})
}

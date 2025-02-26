package utils

import (
    "github.com/gofiber/fiber/v2"
    "log"
)

type ErrorResponse struct {
    Error   string `json:"error"`
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func HandleError(c *fiber.Ctx, err error, status int) error {
    log.Printf("Error: %v", err)
    return c.Status(status).JSON(ErrorResponse{
        Error:   err.Error(),
        Code:    status,
        Message: getErrorMessage(status),
    })
}

func getErrorMessage(status int) string {
    switch status {
    case 429:
        return "Rate limit exceeded"
    case 503:
        return "Service unavailable"
    default:
        return "Internal server error"
    }
}
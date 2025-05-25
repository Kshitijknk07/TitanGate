package middleware

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func ValidateRequest(model interface{}) fiber.Handler {
	validate := validator.New()

	return func(c *fiber.Ctx) error {
		modelType := reflect.TypeOf(model)
		modelInstance := reflect.New(modelType).Interface()

		if err := json.NewDecoder(c.Request().BodyStream()).Decode(modelInstance); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if err := validate.Struct(modelInstance); err != nil {
			var errors []ValidationError
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, ValidationError{
					Field: err.Field(),
					Tag:   err.Tag(),
					Value: fmt.Sprintf("%v", err.Value()),
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Validation failed",
				"errors": errors,
			})
		}

		c.Locals("validatedModel", modelInstance)
		return c.Next()
	}
}

func ValidateQuery(model interface{}) fiber.Handler {
	validate := validator.New()

	return func(c *fiber.Ctx) error {
		modelType := reflect.TypeOf(model)
		modelInstance := reflect.New(modelType).Interface()

		if err := c.QueryParser(modelInstance); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid query parameters",
			})
		}

		if err := validate.Struct(modelInstance); err != nil {
			var errors []ValidationError
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, ValidationError{
					Field: err.Field(),
					Tag:   err.Tag(),
					Value: fmt.Sprintf("%v", err.Value()),
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Validation failed",
				"errors": errors,
			})
		}

		c.Locals("validatedModel", modelInstance)
		return c.Next()
	}
}

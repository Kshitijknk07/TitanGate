package middleware

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/gofiber/fiber/v2"
)

type TransformConfig struct {
	RequestTransform  bool
	ResponseTransform bool
	InputFormat       string // "json", "xml", "yaml"
	OutputFormat      string // "json", "xml", "yaml"
	HeaderTransform   map[string]string
	QueryTransform    map[string]string
}

func NewTransformConfig() TransformConfig {
	return TransformConfig{
		RequestTransform:  true,
		ResponseTransform: true,
		InputFormat:       "json",
		OutputFormat:      "json",
		HeaderTransform:   make(map[string]string),
		QueryTransform:    make(map[string]string),
	}
}

func validateFormat(format string) error {
	format = strings.ToLower(format)
	if format != "json" && format != "xml" && format != "yaml" {
		return fmt.Errorf("unsupported format: %s", format)
	}
	return nil
}

func TransformerMiddleware(config TransformConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Validate formats
		if err := validateFormat(config.InputFormat); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := validateFormat(config.OutputFormat); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Transform request if enabled
		if config.RequestTransform {
			// Transform headers
			for oldKey, newKey := range config.HeaderTransform {
				if value := c.Get(oldKey); value != "" {
					c.Set(newKey, value)
					c.Request().Header.Del(oldKey)
				}
			}

			// Transform query parameters
			for oldKey, newKey := range config.QueryTransform {
				if value := c.Query(oldKey); value != "" {
					c.Request().URI().QueryArgs().Set(newKey, value)
					c.Request().URI().QueryArgs().Del(oldKey)
				}
			}

			// Transform request body
			if len(c.Body()) > 0 {
				var transformedBody interface{}
				var err error

				// Parse input format
				switch strings.ToLower(config.InputFormat) {
				case "json":
					err = json.Unmarshal(c.Body(), &transformedBody)
				case "xml":
					err = xml.Unmarshal(c.Body(), &transformedBody)
				case "yaml":
					err = yaml.Unmarshal(c.Body(), &transformedBody)
				}

				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Invalid request body format",
					})
				}

				// Convert to output format
				var outputBody []byte
				switch strings.ToLower(config.OutputFormat) {
				case "json":
					outputBody, err = json.Marshal(transformedBody)
				case "xml":
					outputBody, err = xml.Marshal(transformedBody)
				case "yaml":
					outputBody, err = yaml.Marshal(transformedBody)
				}

				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to transform request body",
					})
				}

				c.Request().SetBody(outputBody)
			}
		}

		// Set response content type
		c.Response().Header.Set("Content-Type", getContentType(config.OutputFormat))

		// Continue with the request
		err := c.Next()

		// Transform response if enabled
		if config.ResponseTransform && len(c.Response().Body()) > 0 {
			var transformedBody interface{}
			var err error

			// Parse input format
			switch strings.ToLower(config.InputFormat) {
			case "json":
				err = json.Unmarshal(c.Response().Body(), &transformedBody)
			case "xml":
				err = xml.Unmarshal(c.Response().Body(), &transformedBody)
			case "yaml":
				err = yaml.Unmarshal(c.Response().Body(), &transformedBody)
			}

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to parse response body",
				})
			}

			// Convert to output format
			var outputBody []byte
			switch strings.ToLower(config.OutputFormat) {
			case "json":
				outputBody, err = json.Marshal(transformedBody)
			case "xml":
				outputBody, err = xml.Marshal(transformedBody)
			case "yaml":
				outputBody, err = yaml.Marshal(transformedBody)
			}

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to transform response body",
				})
			}

			c.Response().SetBody(outputBody)
		}

		return err
	}
}

func getContentType(format string) string {
	switch strings.ToLower(format) {
	case "json":
		return "application/json"
	case "xml":
		return "application/xml"
	case "yaml":
		return "application/yaml"
	default:
		return "application/json"
	}
}

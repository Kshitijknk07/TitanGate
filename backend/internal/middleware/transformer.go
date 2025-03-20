package middleware

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
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
	DebugMode         bool
}

func NewTransformConfig() TransformConfig {
	return TransformConfig{
		RequestTransform:  true,
		ResponseTransform: true,
		InputFormat:       "json",
		OutputFormat:      "json",
		HeaderTransform:   make(map[string]string),
		QueryTransform:    make(map[string]string),
		DebugMode:         false,
	}
}

func validateFormat(format string) error {
	format = strings.ToLower(format)
	if format != "json" && format != "xml" && format != "yaml" {
		return fmt.Errorf("unsupported format: %s", format)
	}
	return nil
}

func logDebug(config TransformConfig, format string, args ...interface{}) {
	if config.DebugMode {
		log.Printf("[Transformer] "+format, args...)
	}
}

func TransformerMiddleware(config TransformConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Validate formats
		if err := validateFormat(config.InputFormat); err != nil {
			logDebug(config, "Invalid input format: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := validateFormat(config.OutputFormat); err != nil {
			logDebug(config, "Invalid output format: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Transform request if enabled
		if config.RequestTransform {
			logDebug(config, "Processing request transformation")

			// Transform headers
			for oldKey, newKey := range config.HeaderTransform {
				if value := c.Get(oldKey); value != "" {
					logDebug(config, "Transforming header: %s -> %s", oldKey, newKey)
					c.Set(newKey, value)
					c.Request().Header.Del(oldKey)
				}
			}

			// Transform query parameters
			for oldKey, newKey := range config.QueryTransform {
				if value := c.Query(oldKey); value != "" {
					logDebug(config, "Transforming query param: %s -> %s", oldKey, newKey)
					c.Request().URI().QueryArgs().Set(newKey, value)
					c.Request().URI().QueryArgs().Del(oldKey)
				}
			}

			// Transform request body
			if len(c.Body()) > 0 {
				logDebug(config, "Transforming request body from %s to %s", config.InputFormat, config.OutputFormat)
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
					logDebug(config, "Failed to parse request body: %v", err)
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error":   "Invalid request body format",
						"details": err.Error(),
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
					logDebug(config, "Failed to transform request body: %v", err)
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error":   "Failed to transform request body",
						"details": err.Error(),
					})
				}

				c.Request().SetBody(outputBody)
				logDebug(config, "Request body transformation completed")
			}
		}

		// Set response content type
		c.Response().Header.Set("Content-Type", getContentType(config.OutputFormat))

		// Continue with the request
		err := c.Next()

		// Transform response if enabled
		if config.ResponseTransform && len(c.Response().Body()) > 0 {
			logDebug(config, "Processing response transformation")
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
				logDebug(config, "Failed to parse response body: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Failed to parse response body",
					"details": err.Error(),
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
				logDebug(config, "Failed to transform response body: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Failed to transform response body",
					"details": err.Error(),
				})
			}

			c.Response().SetBody(outputBody)
			logDebug(config, "Response body transformation completed")
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

package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthConfig struct {
	SecretKey     string
	TokenExpiry   time.Duration
	ExcludedPaths []string
}

func NewAuthConfig() AuthConfig {
	return AuthConfig{
		SecretKey:   os.Getenv("JWT_SECRET_KEY"),
		TokenExpiry: 24 * time.Hour,
		ExcludedPaths: []string{
			"/health",
			"/metrics",
			"/api/v1/auth/login",
			"/api/v1/auth/register",
		},
	}
}

func AuthMiddleware(config AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()
		for _, excluded := range config.ExcludedPaths {
			if path == excluded {
				return c.Next()
			}
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := validateToken(tokenString, config.SecretKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		c.Locals("user", claims)
		return c.Next()
	}
}

func validateToken(tokenString, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
}

func GenerateToken(userID string, config AuthConfig) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(config.TokenExpiry).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SecretKey))
} 
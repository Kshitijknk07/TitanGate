package handlers

import (
	"encoding/json"
	"time"

	"github.com/Kshitijknk07/TitanGate/backend/internal/cache"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var redisCache = cache.NewRedisCache(redisClient)

func GetUserHandler(c *fiber.Ctx) error {
	userID := c.Query("id")
	cacheKey := "user:" + userID

	cachedData, found := redisCache.Get(cacheKey)
	if found {
		switch v := cachedData.(type) {
		case []byte:
			return c.Type("json").Send(v)
		case string:
			return c.Type("json").Send([]byte(v))
		default:
		}
	}

	userData := map[string]string{"id": userID, "name": "John Doe"}
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to marshal user data")
	}

	redisCache.Set(cacheKey, jsonData, 5*time.Minute)

	return c.JSON(userData)
}

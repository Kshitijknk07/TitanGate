package handlers

import (
    "encoding/json"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/Kshitijknk07/TitanGate/backend/internal/cache"
)

var redisCache = cache.NewRedisCache()

func GetUserHandler(c *fiber.Ctx) error {
    userID := c.Query("id")
    cacheKey := "user:" + userID

    cachedData, err := redisCache.Get(cacheKey)
    if err == nil {
        return c.Type("json").Send([]byte(cachedData))
    }

    userData := map[string]string{"id": userID, "name": "John Doe"}
    jsonData, _ := json.Marshal(userData)
    redisCache.Set(cacheKey, string(jsonData), 5*time.Minute)

    return c.JSON(userData)
}

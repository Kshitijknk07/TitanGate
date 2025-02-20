package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"titangate/internal/cache"
)

var redisCache = cache.NewRedisCache()

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	cacheKey := "user:" + userID

	
	cachedData, err := redisCache.Get(cacheKey)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedData))
		return
	}

	
	userData := map[string]string{"id": userID, "name": "John Doe"}

	
	jsonData, _ := json.Marshal(userData)
	redisCache.Set(cacheKey, string(jsonData), 5*time.Minute)

	
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

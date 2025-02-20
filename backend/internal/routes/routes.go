package routes

import (
	"github.com/Kshitijknk07/TitanGate/backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"	
)


func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")

	return r
}
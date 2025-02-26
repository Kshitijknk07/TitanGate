package routes

import "github.com/gofiber/fiber/v2"

type VersionedRouter struct {
    app *fiber.App
}

func NewVersionedRouter(app *fiber.App) *VersionedRouter {
    return &VersionedRouter{app: app}
}

func (vr *VersionedRouter) Group(version string) fiber.Router {
    return vr.app.Group("/api/" + version)
}
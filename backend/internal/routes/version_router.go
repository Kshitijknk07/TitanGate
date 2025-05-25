package routes

import (
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

const (
	MinVersion     = 1
	CurrentVersion = 2
)

type VersionedRouter struct {
	app               *fiber.App
	supportedVersions map[string]bool
}

func NewVersionedRouter(app *fiber.App) *VersionedRouter {

	versions := make(map[string]bool)
	for v := MinVersion; v <= CurrentVersion; v++ {
		versions[fmt.Sprintf("v%d", v)] = true
	}

	return &VersionedRouter{
		app:               app,
		supportedVersions: versions,
	}
}

func (vr *VersionedRouter) Group(version string) fiber.Router {
	if !vr.isValidVersion(version) {
		panic(fmt.Sprintf("Invalid API version: %s", version))
	}
	return vr.app.Group("/api/" + version)
}

func (vr *VersionedRouter) isValidVersion(version string) bool {

	matched, _ := regexp.MatchString(`^v[1-9]\d*$`, version)
	if !matched {
		return false
	}

	return vr.supportedVersions[version]
}

func (vr *VersionedRouter) GetLatestVersion() string {
	return fmt.Sprintf("v%d", CurrentVersion)
}

func (vr *VersionedRouter) IsVersionSupported(version string) bool {
	return vr.supportedVersions[version]
}

func (vr *VersionedRouter) GetSupportedVersions() []string {
	versions := make([]string, 0, len(vr.supportedVersions))
	for v := range vr.supportedVersions {
		versions = append(versions, v)
	}
	return versions
}

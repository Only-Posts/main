package app

import (
	"github.com/gofiber/fiber"
	"only-post-api/internal/http"
)

func SetupRoutes(app *fiber.App) {
	routes := http.GetRoutes()
	for _, route := range routes {
		app.Get(route.Path, route.Handler)
	}
}

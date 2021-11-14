package app

import (
	"github.com/gofiber/fiber"
	"only-post-api/internal/http"
)

func SetupRoutes(app *fiber.App) {
	routes := http.GetRoutes()
	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler)
	}
}

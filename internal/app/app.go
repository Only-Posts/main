package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"only-post-api/internal/http"
)

func SetupRoutes(app *fiber.App) {
	routes := http.GetRoutes()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		Next:             nil,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTION",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	for _, route := range routes {
		app.Add(route.Method, route.Path, route.Handler)
	}
}

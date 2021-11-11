package main

import (
	"github.com/gofiber/fiber"
	"only-post-api/internal/app"
)

func main() {
	App := fiber.New()
	app.SetupRoutes(App)
	App.Listen(8080)

}

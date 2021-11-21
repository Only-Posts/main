package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"only-post-api/internal/app"
	"only-post-api/internal/database"
)

func main() {
	App := fiber.New()
	err := database.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return
	}

	app.SetupRoutes(App)
	App.Listen(8080)

}

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"only-post-api/internal/app"
	"only-post-api/internal/database"
)

func main() {
	App := fiber.New()
	for {
		err := database.InitDB()
		if err != nil {
			fmt.Println("Failed to connect to db")
		} else {
			break
		}
	}

	app.SetupRoutes(App)
	log.Println(App.Listen(":8080"))

}

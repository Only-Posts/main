package endpoint

import "github.com/gofiber/fiber/v2"

func Main(ctx *fiber.Ctx) error {
	ctx.SendString("HELLO WORLD!")
	return nil
}

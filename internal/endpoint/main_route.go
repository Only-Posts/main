package endpoint

import "github.com/gofiber/fiber"

func Main(ctx *fiber.Ctx) {
	ctx.Send("HELLO WORLD!")
}

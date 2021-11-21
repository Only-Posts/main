package endpoint

import "github.com/gofiber/fiber"

func PostLogin(ctx *fiber.Ctx) {
	ctx.Send("HELLO LOGIN!")
}

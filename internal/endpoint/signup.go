package endpoint

import (
	"github.com/gofiber/fiber"
	"only-post-api/internal/auth"
)

func PostSignUp(ctx *fiber.Ctx) {
	var user *auth.UserData
	err := ctx.BodyParser(user)
	if err != nil {
		ctx.JSON(fiber.Map{"status": "failed", "message": "failed to parse body"})
		return
	}
	ctx.JSON(fiber.Map{"status": "success", "message": "user registered"})
}

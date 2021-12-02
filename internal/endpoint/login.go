package endpoint

import (
	"log"

	"github.com/gofiber/fiber"
	"only-post-api/internal/auth"
	"only-post-api/internal/repository"
)

func PostLogin(ctx *fiber.Ctx) {
	user := new(auth.UserData)
	err := ctx.BodyParser(user)
	if err != nil {
		ctx.JSON(fiber.Map{"status": "failed", "message": "failed to parse body"})
		log.Printf("Error parse signup message err: %s\n", err)
		return
	}

	getuser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return
	}
	if getuser.Email != "" {
		if getuser.Password != user.Password {
			ctx.JSON(fiber.Map{"status": "failed", "message": "wrong password"})
			return
		}
		ctx.JSON(fiber.Map{"status": "success", "message": "user login", "user_id": getuser.ID})
		return
	}

	getuser, err = repository.GetUserByUsername(user.Username)
	if err != nil {
		return
	}
	if getuser.Username != "" {
		if getuser.Password != user.Password {
			ctx.JSON(fiber.Map{"status": "failed", "message": "wrong password"})
			return
		}
		ctx.JSON(fiber.Map{"status": "success", "message": "user login", "user_id": getuser.ID})
		return
	}

	ctx.JSON(fiber.Map{"status": "failed", "message": "user not found"})
	return
}

package endpoint

import (
	"log"

	"github.com/gofiber/fiber"
	"only-post-api/internal/auth"
	"only-post-api/internal/repository"
)

func PostSignUp(ctx *fiber.Ctx) {
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
		ctx.JSON(fiber.Map{"status": "success", "message": "user already registered"})
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
		ctx.JSON(fiber.Map{"status": "success", "message": "user already registered"})
		return
	}

	err = repository.RegisterUser(user)
	if err != nil {
		ctx.JSON(fiber.Map{"status": "failed", "message": "failed to register"})
		return
	}
	ctx.JSON(fiber.Map{"status": "success", "message": "user registered", "user_id": user.ID})
	return
}

package endpoint

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"only-post-api/internal/auth"
	"only-post-api/internal/repository"
)

func PostSignUp(ctx *fiber.Ctx) error {
	user := new(auth.UserData)
	err := ctx.BodyParser(user)
	if err != nil {
		ctx.JSON(fiber.Map{"status": "failed", "message": "failed to parse body"})
		log.Printf("Error parse signup message err: %s\n", err)

	}

	getuser, err := repository.GetUserByEmail(user.Email)
	if err != nil {

	}
	if getuser.Email != "" {
		if getuser.Password != user.Password {
			ctx.JSON(fiber.Map{"status": "failed", "message": "wrong password"})

		}
		ctx.JSON(fiber.Map{"status": "success", "message": "user already registered"})
		return nil
	}

	getuser, err = repository.GetUserByUsername(user.Username)
	if err != nil {

	}
	if getuser.Username != "" {
		if getuser.Password != user.Password {
			ctx.JSON(fiber.Map{"status": "failed", "message": "wrong password"})

		}
		ctx.JSON(fiber.Map{"status": "success", "message": "user already registered"})
		return nil
	}

	err = repository.RegisterUser(user)
	if err != nil {
		ctx.JSON(fiber.Map{"status": "failed", "message": "failed to register"})

	}
	ctx.JSON(fiber.Map{"status": "success", "message": "user registered", "user_id": user.ID})
	return nil
}

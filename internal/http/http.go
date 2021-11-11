package http

import (
	"github.com/gofiber/fiber"
)

const (
	Get  = "GET"
	Post = "POST"
)

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

func GetRoutes() []Route {
	return []Route{
		{
			Method:  Get,
			Path:    "/",
			Handler: func(ctx *fiber.Ctx) {},
		},
	}
}

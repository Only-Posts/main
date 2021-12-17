package http

import (
	"github.com/gofiber/fiber/v2"
	"only-post-api/internal/endpoint"
)

const (
	get  = "GET"
	post = "POST"
)

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

func GetRoutes() []Route {
	return []Route{
		{
			Method:  post,
			Path:    "/user/login",
			Handler: endpoint.PostLogin,
		},
		{
			Method:  post,
			Path:    "/user/signup",
			Handler: endpoint.PostSignUp,
		},
		{
			Method:  get,
			Path:    "/",
			Handler: endpoint.Main,
		},
	}
}

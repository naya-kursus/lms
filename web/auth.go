package web

import (
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
	"github.com/hadihammurabi/belajar-go-rest-api/web/dto"

	"github.com/gofiber/fiber/v2"
)

type auth struct {
	*Web
	router *fiber.App

	service *service.Service
}

// Auth func
func Auth(r *Web) auth {
	api := auth{r, fiber.New(), service.Get()}
	api.router.Post("/login", api.Login)

	return api
}

// Login func
func (api auth) Login(c *fiber.Ctx) error {
	input := &dto.UserLoginRequest{}
	if err := c.BodyParser(input); err != nil {
		return dto.Fail(c, err)
	}

	user := &entity.User{
		Email:    input.Email,
		Password: input.Password,
	}

	token, err := api.service.Auth.Login(c.Context(), user)
	if err != nil {
		return dto.Fail(c, "invalid credentials")
	}

	return dto.Ok(c, &dto.UserLoginResponse{
		Token: token,
		Type:  "Bearer",
	})
}

// Me func
func (api auth) Me(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*entity.User)
	return dto.Ok(c, fromLocals)
}

package web

import (
	"github.com/gofiber/fiber/v2"
)

type index struct {
	*Web
	router *fiber.App
}

func Index(r *Web) *fiber.App {
	api := index{r, fiber.New()}
	api.router.Get("", api.Index)

	return api.router
}

func (api index) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Selamat datang di Belajar REST API dengan Go",
	})
}

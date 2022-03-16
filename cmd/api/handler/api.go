package handler

import (
	"github.com/gofiber/fiber/v2"
)

// HelloWorld godoc
// @Summary Get a Hello World!
// @Param name path string Hello who?
// @Success 200
// @Router /hello/{name?} [get]
func (this *Handler) HelloWorld(c *fiber.Ctx) (err error) {
	name := c.Params("name")
	if name == "" {
		return c.SendString("Hello, World!")
	}
	return c.SendString("Hello, " + name)
}

package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Template godoc
// @Summary Get a simple template
// @Success 200
// @Router /template [get]
func (this *Handler) Template(c *fiber.Ctx) (err error) {
	return c.Render("index", fiber.Map{
		"Title": "Go Serverless!",
	})
}

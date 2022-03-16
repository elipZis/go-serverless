package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Handler stores the references to services and repositories for controllers to use
type Handler struct {
	app *fiber.App
}

// NewHandler returns a new app handler
func NewHandler(app *fiber.App) (this *Handler) {
	this = new(Handler)
	this.app = app
	return this
}

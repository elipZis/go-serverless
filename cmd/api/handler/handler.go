package handler

import (
	"github.com/elipzis/go-serverless/api/repository"
	"github.com/gofiber/fiber/v2"
)

// Handler stores the references to services and repositories for controllers to use
type Handler struct {
	app  *fiber.App
	repo *repository.Repository
}

// NewHandler returns a new app handler
func NewHandler(app *fiber.App) (this *Handler) {
	this = new(Handler)
	this.app = app
	this.repo = repository.NewRepository()
	return this
}

// SendError creates a status response of the given error
func SendError(c *fiber.Ctx, status int, err error) error {
	return SendErrorFromString(c, status, err.Error())
}

// SendErrorFromString creates a status response of the given error string
func SendErrorFromString(c *fiber.Ctx, status int, err string) error {
	return c.Status(status).SendString(err)
}

// bindRequest is a generic binding of requests to models
func (this *Handler) bindRequest(c *fiber.Ctx, model interface{}) error {
	if err := c.BodyParser(model); err != nil {
		return err
	}
	return nil
}

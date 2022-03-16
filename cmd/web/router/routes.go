package router

import (
	"github.com/elipzis/go-serverless/web/handler"
	"github.com/gofiber/fiber/v2"
)

// Register the routes
func (this *Router) Register(rootGroup fiber.Router) {
	// Handler
	controller := handler.NewHandler(this.App)

	// Handlers
	rootGroup.Get("/", controller.Template) // /
}

package router

import (
	"github.com/elipzis/go-serverless/api/handler"
	"github.com/elipzis/go-serverless/api/router/middleware"
	"github.com/gofiber/fiber/v2"
)

// Register the routes
func (this *Router) Register(rootGroup fiber.Router) {
	// Handler
	controller := handler.NewHandler(this.App)

	// Handlers
	rootGroup.Get("/hello/:name?", controller.HelloWorld)                                // /hello/{name?}
	rootGroup.Get("/secret/hello/:name?", middleware.Protected(), controller.HelloWorld) // /secret/hello/{name?}
	rootGroup.Get("/queue/:name?", controller.QueueHelloWorld)                           // /queue/{name?}
}

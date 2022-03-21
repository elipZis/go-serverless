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

	// Deployed subgroup
	v1 := rootGroup.Group("/api") // /api

	// Handlers
	v1.Get("/hello/:name?", controller.HelloWorld)                                // /hello/{name?}
	v1.Get("/secret/hello/:name?", middleware.Protected(), controller.HelloWorld) // /secret/hello/{name?}
	v1.Get("/queue/:name?", controller.QueueHelloWorld)                           // /queue/{name?}
}

package router

import (
	"context"
	rice "github.com/GeertJohan/go.rice"
	"github.com/elipzis/go-serverless/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

type Router struct {
	App *fiber.App
}

// NewRouter Create a new router and configure some middlewares
func NewRouter() (this *Router) {
	this = new(Router)

	// Create a template filesystem of compiled templates
	engine := html.NewFileSystem(rice.MustFindBox("../view").HTTPBox(), ".html")
	this.App = fiber.New(fiber.Config{
		AppName: util.GetEnvOrDefault("APP_NAME", "Go Serverless!"),
		Views:   engine,
	})

	// Global middlewares
	this.App.Use(cors.New())
	this.App.Use(logger.New())
	this.App.Use(compress.New())

	return this
}

// Run Start the server
func (this *Router) Run() {
	// Start server
	go func() {
		log.Printf("[Router] Trying to start a server at %s:%s", util.GetEnvOrDefault("HOST", ""), util.GetEnvOrDefault("PORT", "1323"))
		if err := this.App.Listen(util.GetEnvOrDefault("HOST", "") + ":" + util.GetEnvOrDefault("PORT", "1323")); err != nil {
			log.Print("[Router] Shutting down the server...")
		}
	}()
}

// Shutdown calm and correctly or fatal
func (this *Router) Shutdown(ctx context.Context) {
	if err := this.App.Shutdown(); err != nil {
		log.Fatal(err)
	}
}

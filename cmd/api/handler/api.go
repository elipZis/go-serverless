package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
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

// QueueHelloWorld godoc
// @Summary Send a Hello World! to a queue
// @Param name path string Hello who?
// @Success 200
// @Router /queue/{name?} [get]
func (this *Handler) QueueHelloWorld(c *fiber.Ctx) (err error) {
	str := "Hello, World!"
	name := c.Params("name")
	if name != "" {
		str = "Hello, " + name
	}

	// Create a simple queue message
	msg, err := json.Marshal(map[string]string{
		"name": str,
	})
	if err != nil {
		log.Println("[Enrichment] Error while marshalling", err)
	}
	_, err = this.repo.SendMessage(msg)
	if err != nil {
		log.Println("[Enrichment] Error while pushing into the queue", err)
	}
	return err
}

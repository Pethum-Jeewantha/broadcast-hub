package app

import (
	"broadcast-hub/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Start() error {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(handler.WebsocketHandler))

	return app.Listen(":3200")
}

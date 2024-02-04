package app

import (
	"broadcast-hub/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Start() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(handler.WebsocketHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3200" // Default port value if PORT is not set
	}

	return app.Listen(":" + port)
}

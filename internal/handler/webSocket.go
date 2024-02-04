package handler

import (
	"broadcast-hub/internal/hub"
	"github.com/gofiber/websocket/v2"
)

func WebsocketHandler(c *websocket.Conn) {
	hub.RegisterClient(c)

	defer hub.UnregisterClient(c)

	for {
		if _, message, err := c.ReadMessage(); err != nil {
			break
		} else {
			hub.Broadcast(message)
		}
	}
}

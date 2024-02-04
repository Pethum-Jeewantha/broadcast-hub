package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	ws "github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = ws.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins
		},
	}

	clients = make(map[*websocket.Conn]bool)
	mu      sync.Mutex
)

func broadcast(message []byte) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("error: %v", err)
			err := client.Close()
			if err != nil {
				return
			}
			delete(clients, client)
		}
	}
}

func wsHandler(c *websocket.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, c)
		mu.Unlock()
		err := c.Close()
		if err != nil {
			return
		}
	}()

	mu.Lock()
	clients[c] = true
	mu.Unlock()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		broadcast(message)
	}
}

func main() {
	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the request is a WebSocket upgrade
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		wsHandler(c)
	}))

	log.Fatal(app.Listen(":3200"))
}

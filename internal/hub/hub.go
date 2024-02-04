package hub

import (
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

var (
	clients = make(map[*websocket.Conn]bool)
	mu      sync.Mutex
)

func Broadcast(message []byte) {
	mu.Lock()
	defer mu.Unlock()

	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("error: %v", err)
			err := client.Close()
			if err != nil {
				return
			}
			delete(clients, client)
		}
	}
}

func RegisterClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	clients[client] = true
}

func UnregisterClient(client *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	delete(clients, client)
	err := client.Close()
	if err != nil {
		return
	}
}

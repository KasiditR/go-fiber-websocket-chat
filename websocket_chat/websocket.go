package websocket_chat

import (
	"log"

	"github.com/KasiditR/websocket/message"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type RegisterChan chan *websocket.Conn
type UnregisterChan chan *websocket.Conn
type BroadcastChan chan message.Message

func SetupWebSocketRoutes(app *fiber.App, register RegisterChan, unregister UnregisterChan, broadcast BroadcastChan) {
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		register <- c
		defer func() { unregister <- c }()

		for {
			var msg message.Message
			if err := c.ReadJSON(&msg); err != nil {
				log.Println("Error reading message:", err)
				break
			}
			broadcast <- msg
		}
	}))
}

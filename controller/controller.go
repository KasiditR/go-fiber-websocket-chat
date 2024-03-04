package controller

import (
	"github.com/KasiditR/websocket/client"
	"github.com/KasiditR/websocket/websocket_chat"
	"github.com/gofiber/fiber/v2"
	"log"
)

func SetupController(app *fiber.App) {
	register := make(websocket_chat.RegisterChan)     // Channel for registering new clients
	unregister := make(websocket_chat.UnregisterChan) // Channel for unregistering clients
	broadcast := make(websocket_chat.BroadcastChan)   // Channel for broadcasting messages

	go client.HandleClients(register, unregister)
	go BroadcastMessages(broadcast)
	websocket_chat.SetupWebSocketRoutes(app, register, unregister, broadcast)
}

func BroadcastMessages(broadcast websocket_chat.BroadcastChan) {
	for {
		msg := <-broadcast
		for conn, clientValue := range client.Clients {
			if msg.Type == "public" || (msg.Type == "private" && clientValue.ID == msg.Target) {
				err := conn.WriteJSON(map[string]interface{}{
					"sender":  clientValue.ID,
					"content": msg.Content,
					"type":    msg.Type,
				})
				if err != nil {
					log.Println("Error sending message:", err)
					conn.Close()
					delete(client.Clients, conn)
				}
			}
		}
	}
}

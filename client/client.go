package client

import (
	"fmt"
	"github.com/KasiditR/websocket/websocket_chat"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:   id,
		Conn: conn,
	}
}

var Clients = make(map[*websocket.Conn]*Client) // Connected clients

func HandleClients(register websocket_chat.RegisterChan, unregister websocket_chat.UnregisterChan) {
	for {
		select {
		case conn := <-register:
			// Register new client
			clientID := conn.Params("id")
			newClient := NewClient(clientID, conn)
			Clients[conn] = newClient
			fmt.Println("Client connected:", clientID)
		case conn := <-unregister:
			// Remove client on disconnect
			client := Clients[conn]
			if client != nil {
				delete(Clients, conn)
				fmt.Println("Client disconnected:", client.ID)
			}
		}
	}
}
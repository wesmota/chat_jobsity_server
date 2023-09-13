package websocket

import (
	"github.com/rs/zerolog/log"

	"github.com/wesmota/go-jobsity-chat-server/models"
)

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan models.ChatMessage

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan models.ChatMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client] = true
			log.Info().Msgf("Client registered: %s", client.ChatUser.Email)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				log.Info().Msgf("Client unregistered: %s", client.ChatUser.Email)
			}

		case msg := <-h.Broadcast:
			for c := range h.Clients {
				err := c.Connection.WriteJSON(msg)
				log.Info().Msgf("Broadcasting message: %+v", msg)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

package websocket

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/models"
	"github.com/wesmota/go-jobsity-chat-server/usecase"
)

type ChatUser struct {
	Email string
}

// Client is a websocket client, basically a frontend visitor
type Client struct {
	// the websocket connection
	Connection      *websocket.Conn
	ChatUser        ChatUser
	Hub             *Hub
	ChatRoomService *usecase.ChatRoomService
}

// NewClient is used to initialize a new Client with all required values initialized
func NewClient(conn *websocket.Conn, user ChatUser, hub *Hub, chatService *usecase.ChatRoomService) *Client {
	return &Client{
		Connection:      conn,
		ChatUser:        user,
		Hub:             hub,
		ChatRoomService: chatService,
	}
}

func (c *Client) Read(msgChan chan []byte) {
	for {
		messageType, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Err(err).Msg("error on reading client message :")
			return
		}
		var chatMsg models.ChatMessage
		err = json.Unmarshal(p, &chatMsg)
		if err != nil {
			log.Err(err).Msg("error on unmarshalling message read:")
			return
		}
		chatMsg.ChatUser = c.ChatUser.Email
		chatMsg.Type = messageType
		c.Hub.Broadcast <- chatMsg
		log.Info().Interface("chatMsg", chatMsg).Msg("Read")
		msgChan <- p
		go c.ChatRoomService.CreateChatMessage(context.Background(), chatMsg)

	}
}

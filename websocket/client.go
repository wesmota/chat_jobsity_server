package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/wesmota/go-jobsity-chat-server/models"
)

type ChatUser struct {
	ID    uint
	Email string
}

// Client is a websocket client, basically a frontend visitor
type Client struct {
	// the websocket connection
	Connection *websocket.Conn
	ChatUser   ChatUser
	Hub        *Hub
}

// NewClient is used to initialize a new Client with all required values initialized
func NewClient(conn *websocket.Conn, user ChatUser) *Client {
	return &Client{
		Connection: conn,
		ChatUser:   user,
	}
}

func (c *Client) Read(msgChan chan []byte) {
	for {
		messageType, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println("error:", err)
			return
		}
		var chatMsg models.ChatMessage
		err = json.Unmarshal(p, &chatMsg)
		if err != nil {
			log.Println("error:", err)
			return
		}
		chatMsg.ChatUser = c.ChatUser.Email
		chatMsg.Type = messageType
		c.Hub.Broadcast <- chatMsg
		log.Println("info:", "Message received: ", chatMsg)
		msgChan <- p
	}
}

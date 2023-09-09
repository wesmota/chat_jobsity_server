package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	chatwebsocket "github.com/wesmota/go-jobsity-chat-server/websocket"
)

var (
	/**
	websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
	*/
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	hub := chatwebsocket.NewHub()
	go hub.Run()

	client := &chatwebsocket.Client{
		Connection: conn,
		Hub:        hub,
		ChatUser: chatwebsocket.ChatUser{
			ID:    1,
			Email: "teste@gmail.com",
		},
	}

	hub.Register <- client
	requestBody := make(chan []byte)
	go client.Read(requestBody)
}

package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
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
	log.Info().Msg("ServeWS")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	log.Info().Msg("New websocket connection")

	jwtToken := r.URL.Query().Get("jwt")
	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		log.Err(err)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Err(err)
		return
	}
	log.Info().Interface("claims", claims).Msg("JWT Claims")

	hub := chatwebsocket.NewHub()
	go hub.Run()

	chatUser := chatwebsocket.ChatUser{
		Email: claims["Email"].(string),
	}
	log.Info().Interface("chatUser", chatUser).Msg("ServeWS")
	client := chatwebsocket.NewClient(conn, chatUser, hub, h.ChatRoomService)
	log.Info().Interface("client", client).Msg("ServeWS")

	hub.Register <- client
	requestBody := make(chan []byte)
	go client.Read(requestBody)
	go h.Broker.Read(hub)
	go h.Broker.Publish(requestBody)

}

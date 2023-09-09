package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/wesmota/go-jobsity-chat-server/handlers"
	"github.com/wesmota/go-jobsity-chat-server/logger"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := logger.NewZerologLogger(zerolog.New(os.Stdout))
	ctx := context.Background()
	h := handlers.NewDefaultHandler(ctx)
	// Setup app routes
	r := mux.NewRouter()
	sb := r.PathPrefix("/v1").Subrouter()
	sb.HandleFunc("/ws", h.ServeWS)
	sb.HandleFunc("/rooms", h.ListChatRooms).Methods("GET")
	sb.HandleFunc("/rooms", h.CreateChatRoom).Methods("POST")
	sb.HandleFunc("/rooms/{room_id}/messages", h.CreateChatMessage).Methods("POST")
	sb.HandleFunc("/users", h.CreateUser).Methods("POST")
	// Start api server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal().Err(http.ListenAndServe(":"+port, r)).Msg("Server error")
	log.Info().Msgf("Started server on port %s", port)

}

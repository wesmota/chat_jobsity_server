package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/rs/zerolog"
	"github.com/wesmota/go-jobsity-chat-server/handlers"
	"github.com/wesmota/go-jobsity-chat-server/handlers/middlewares"
	"github.com/wesmota/go-jobsity-chat-server/logger"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := logger.NewZerologLogger(zerolog.New(os.Stdout))
	ctx := context.Background()
	h := handlers.NewDefaultHandler(ctx)

	// Setup app routes
	r := mux.NewRouter()
	r.Use(middlewares.Logger)
	sb := r.PathPrefix("/v1").Subrouter()
	sb.HandleFunc("/api/auth/signup", h.SignUp).Methods("POST")
	sb.HandleFunc("/api/auth/login", h.Login).Methods("POST")
	sb.HandleFunc("/ws", h.ServeWS)

	sbChat := r.PathPrefix("/v1/api/chat").Subrouter()
	sbChat.Use(middlewares.Authenticate)
	sbChat.HandleFunc("/rooms", h.ListChatRooms).Methods("GET")
	sbChat.HandleFunc("/rooms", h.CreateChatRoom).Methods("POST")
	sbChat.HandleFunc("/rooms/{room_id}/messages", h.CreateChatMessage).Methods("POST")

	// Start api server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal().Err(http.ListenAndServe(":"+port, r)).Msg("Server error")
	log.Info().Msgf("Started server on port %s", port)

}

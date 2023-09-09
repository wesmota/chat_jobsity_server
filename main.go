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
	// Start api server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Info().Msgf("Starting server on port %s", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, r)).Msg("Server error")

}

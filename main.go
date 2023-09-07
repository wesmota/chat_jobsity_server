package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Setup app routes
	r := mux.NewRouter()

}

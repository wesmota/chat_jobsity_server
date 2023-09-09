package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/db"
	"github.com/wesmota/go-jobsity-chat-server/handlers/presenter"
	"github.com/wesmota/go-jobsity-chat-server/logger"
	chatrooms "github.com/wesmota/go-jobsity-chat-server/storage/chat_rooms"
	usecase "github.com/wesmota/go-jobsity-chat-server/usecase"
)

var (
	ErrInRequestMarshaling = errors.New("invalid/bad request paramenters")
)

// Handler represents an HTTP request Handler.
type Handler struct {
	ChatRoomService *usecase.ChatRoomService
}

func NewDefaultHandler(ctx context.Context) *Handler {
	log := logger.NewZerologLogger(zerolog.New(os.Stdout))
	database, err := db.NewDB()
	if err != nil {
		log.Err(err).Msg("DB Connection error, initial format")
	}
	chatsRepo, err := chatrooms.NewChatRoomsRepo(database, log)
	if err != nil {
		log.Err(err).Msg("DB Connection error, initial format")
	}
	chatRoomService := usecase.New(chatsRepo)
	return &Handler{
		ChatRoomService: chatRoomService,
	}
}

func ErrResponse(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errCode := codeFrom(err)
	w.WriteHeader(errCode)
	res := presenter.ErrorResponse{Message: err.Error(), Status: false, Code: errCode}
	data, err := json.Marshal(res)
	w.Write(data)
}

// codeFrom returns the http status code from service errors
func codeFrom(err error) int {
	switch err {
	case ErrInRequestMarshaling:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func (h *Handler) ListChatRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rooms, err := h.ChatRoomService.ListChatRooms(context.Background())
	log.Info().Interface("rooms", rooms).Msg("ListChatRooms")
	if err != nil {
		//ErrResponse(err, w)
		return
	}
	// transform rooms to json
	_, err = json.Marshal(rooms)
	if err != nil {
		log.Info().Msgf("Error marshaling rooms: %v", err)
		//ErrResponse(err, w)
		return
	}
	json.NewEncoder(w).Encode(rooms)
	log.Info().Msg("ListChatRooms handler concluded")
	return
}

func (h *Handler) CreateChatRoom(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.ChatRoomService.ListChatRooms(context.Background())
	if err != nil {
		ErrResponse(err, w)
		return
	}
	data, err := json.Marshal(rooms)
	if err != nil {
		ErrResponse(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

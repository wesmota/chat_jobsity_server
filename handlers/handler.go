package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wesmota/go-jobsity-chat-server/db"
	"github.com/wesmota/go-jobsity-chat-server/handlers/presenter"
	"github.com/wesmota/go-jobsity-chat-server/logger"
	"github.com/wesmota/go-jobsity-chat-server/models"
	"github.com/wesmota/go-jobsity-chat-server/rabbitmq"
	"github.com/wesmota/go-jobsity-chat-server/storage/authorization"
	chatrooms "github.com/wesmota/go-jobsity-chat-server/storage/chat_rooms"
	usecase "github.com/wesmota/go-jobsity-chat-server/usecase"
)

var (
	ErrInRequestMarshaling = errors.New("invalid/bad request paramenters")
)

// Handler represents an HTTP request Handler.
type Handler struct {
	ChatRoomService *usecase.ChatRoomService
	AuthService     *usecase.AuthService
	Broker          *rabbitmq.Broker
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
	chatRoomService := usecase.NewChatService(chatsRepo)

	authRepo, err := authorization.NewAuthorizationsRepo(database, log)
	if err != nil {
		log.Err(err).Msg("DB Connection error, initial format")
	}
	authService := usecase.NewAuthService(authRepo)

	// RabbitMQ
	rmqHost := os.Getenv("RMQ_HOST")
	rmqUserName := os.Getenv("RMQ_USERNAME")
	rmqPassword := os.Getenv("RMQ_PASSWORD")
	rmqPort := os.Getenv("RMQ_PORT")
	dsn := "amqp://" + rmqUserName + ":" + rmqPassword + "@" + rmqHost + ":" + rmqPort + "/"

	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	br := &rabbitmq.Broker{}
	br.Setup(ch)

	return &Handler{
		ChatRoomService: chatRoomService,
		AuthService:     authService,
		Broker:          br,
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
	if err != nil {
		ErrResponse(err, w)
		return
	}
	_, err = json.Marshal(rooms)
	if err != nil {
		log.Info().Msgf("Error marshaling rooms: %v", err)
		ErrResponse(err, w)
		return
	}
	json.NewEncoder(w).Encode(rooms)
	log.Info().Msg("ListChatRooms handler concluded")
	return
}

func (h *Handler) CreateChatRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chatRoom models.ChatRoom
	err := json.NewDecoder(r.Body).Decode(&chatRoom)
	if err != nil {
		log.Info().Msgf("Error decoding chat room: %v", err)
		ErrResponse(ErrInRequestMarshaling, w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ChatRoomService.CreateChatRoom(context.Background(), chatRoom)
	if err != nil {
		ErrResponse(err, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Info().Msg("CreateChatRoom handler concluded")
	w.WriteHeader(http.StatusCreated)
	return

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Info().Msgf("Error decoding: %v", err)
		ErrResponse(ErrInRequestMarshaling, w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ChatRoomService.CreateUser(context.Background(), user)
	if err != nil {
		ErrResponse(err, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) CreateChatMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chatMessage models.ChatMessage
	err := json.NewDecoder(r.Body).Decode(&chatMessage)
	if err != nil {
		log.Info().Msgf("Error decoding chat room: %v", err)
		ErrResponse(ErrInRequestMarshaling, w)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get chat user
	err = h.ChatRoomService.CreateChatMessage(context.Background(), chatMessage)
	if err != nil {
		ErrResponse(err, w)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	return
}

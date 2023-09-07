package presenter

import "github.com/wesmota/go-jobsity-chat-server/models"

type ErrorResponse struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
	Status  bool   `json:"Status"`
}

type ChatRoomsResponse struct {
	ChatRooms []models.ChatRoom `json:"ChatRooms"`
}

type ChatRoomResponse struct {
	ChatRoom models.ChatRoom `json:"ChatRoom"`
}

type ChatMessage struct {
	ChatMessage  string `json:"chatMessage"`
	ChatUser     string `json:"chatUser"`
	ChatRoomId   uint   `json:"chatRoomId"`
	ChatRoomName string `json:"chatRoomName"`
}

type ChatRoomMessagesResponse struct {
	Chats []ChatMessage `json:"Chats"`
}

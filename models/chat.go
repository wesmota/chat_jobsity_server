package models

type Chat struct {
	Message    string
	UserId     uint
	User       User
	ChatRoomId uint
	ChatRoom   ChatRoom
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

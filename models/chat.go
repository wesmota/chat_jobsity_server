package models

type Chat struct {
	Message    string
	UserId     uint
	User       User
	ChatRoomId uint
	ChatRoom   ChatRoom
}

type ChatMessage struct {
	Type         int    `json:"type"`
	ChatMessage  string `json:"chatmessage"`
	ChatUser     string `json:"chatuser"`
	ChatRoomId   uint   `json:"chatroomId"`
	ChatRoomName string `json:"chatroomname"`
}

type ChatRoomMessagesResponse struct {
	Chats []ChatMessage `json:"chats"`
}

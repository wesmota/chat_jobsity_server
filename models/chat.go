package models

type ChatMessage struct {
	Type        int    `json:"type"`
	ChatMessage string `json:"chatmessage"`
	ChatUser    string `json:"chatuser"`
	ChatRoomId  uint   `json:"chatroomId"`
}

type ChatRoomMessagesResponse struct {
	Chats []ChatMessage `json:"chats"`
}

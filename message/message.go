package message

import "time"

type Message struct {
	ID             string    `json:"id"  gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ConversationID string    `json:"conversation_id"`
	AuthorID       string    `json:"author_id"`
	Text           string    `json:"text"`
	Timestamp      time.Time `json:"timestamp"`
}

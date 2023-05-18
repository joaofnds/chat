package conversation

import (
	"app/message"
	"app/user"
)

type Conversation struct {
	ID       string            `json:"id"  gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Users    []user.User       `json:"users" gorm:"many2many:conversation_users"`
	Messages []message.Message `json:"messages"  gorm:"foreignKey:ConversationID"`
}

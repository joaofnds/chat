package conversation

import (
	"app/message"
	"app/user"
)

type Conversation struct {
	ID           string            `json:"id"  gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Participants []user.User       `json:"participants" gorm:"many2many:conversation_participants;"`
	Messages     []message.Message `json:"messages"`
}

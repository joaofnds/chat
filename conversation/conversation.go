package conversation

import (
	"app/message"
	"app/user"
	"context"
)

type Conversation struct {
	ID       string            `json:"id"  gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Users    []user.User       `json:"users" gorm:"many2many:conversation_users"`
	Messages []message.Message `json:"messages"  gorm:"foreignKey:ConversationID"`
}

type Repository interface {
	Create(ctx context.Context, user1, user2 user.User) (Conversation, error)
	Find(ctx context.Context, id string) (Conversation, error)
	FindForUser(ctx context.Context, user user.User) ([]Conversation, error)
	AddMessage(ctx context.Context, convo Conversation, msg *message.Message) error
}

type MessagePublisher interface {
	PublishMessage(convo Conversation, author user.User, msg message.Message) error
}

package conversation

import (
	"context"

	"app/message"
	"app/user"

	"gorm.io/gorm"
)

type Service struct {
	orm *gorm.DB
}

func NewService(orm *gorm.DB) *Service {
	return &Service{orm: orm}
}

func (service *Service) Create(ctx context.Context, user1, user2 user.User) (Conversation, error) {
	convo := Conversation{Users: []user.User{user1, user2}, Messages: []message.Message{}}
	return convo, service.orm.WithContext(ctx).Create(&convo).Error
}

func (service *Service) Find(ctx context.Context, id string) (Conversation, error) {
	var convo Conversation

	result := service.orm.
		WithContext(ctx).
		Model(&Conversation{}).
		Preload("Messages").
		Preload("Users").
		Joins("INNER JOIN conversation_users ON conversation_users.conversation_id = conversations.id").
		Joins("INNER JOIN users ON conversation_users.user_id = users.id").
		Find(&convo, "conversations.id = ?", id)

	return convo, result.Error
}

func (service *Service) FindForUser(ctx context.Context, user user.User) ([]Conversation, error) {
	var conversations []Conversation

	result := service.orm.
		WithContext(ctx).
		Model(&Conversation{}).
		Preload("Messages").
		Preload("Users").
		Joins("INNER JOIN conversation_users ON conversation_users.conversation_id = conversations.id").
		Joins("INNER JOIN users ON conversation_users.user_id = users.id").
		Where("users.id = ?", user.ID).
		Find(&conversations)

	return conversations, result.Error
}

func (service *Service) SendMessage(ctx context.Context, convo Conversation, author user.User, text string) (message.Message, error) {
	msg := message.Message{AuthorID: author.ID, Text: text}
	return msg, service.orm.WithContext(ctx).Model(&convo).Association("Messages").Append(&msg)
}

func (service *Service) DeleteAll(ctx context.Context) error {
	return service.orm.WithContext(ctx).Exec("DELETE FROM conversations").Error
}

package conversation

import (
	"app/message"
	"app/user"
	"context"
	"errors"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	orm *gorm.DB
}

func NewPostgresRepository(orm *gorm.DB) *PostgresRepository {
	return &PostgresRepository{orm: orm}
}

func (repository *PostgresRepository) Create(ctx context.Context, user1, user2 user.User) (Conversation, error) {
	convo := Conversation{Users: []user.User{user1, user2}, Messages: []message.Message{}}
	return convo, gormResult(repository.orm.WithContext(ctx).Create(&convo))
}

func (repository *PostgresRepository) Find(ctx context.Context, id string) (Conversation, error) {
	var convo Conversation

	result := repository.orm.
		WithContext(ctx).
		Model(&Conversation{}).
		Preload("Messages").
		Preload("Users").
		Joins("INNER JOIN conversation_users ON conversation_users.conversation_id = conversations.id").
		Joins("INNER JOIN users ON conversation_users.user_id = users.id").
		Find(&convo, "conversations.id = ?", id)

	return convo, gormResult(result)
}

func (repository *PostgresRepository) FindForUser(ctx context.Context, user user.User) ([]Conversation, error) {
	var conversations []Conversation

	result := repository.orm.
		WithContext(ctx).
		Model(&Conversation{}).
		Preload("Messages").
		Preload("Users").
		Joins("INNER JOIN conversation_users ON conversation_users.conversation_id = conversations.id").
		Joins("INNER JOIN users ON conversation_users.user_id = users.id").
		Where("users.id = ?", user.ID).
		Find(&conversations)

	return conversations, gormResult(result)
}

func (repository *PostgresRepository) AddMessage(ctx context.Context, convo Conversation, msg *message.Message) error {
	err := repository.orm.WithContext(ctx).Model(&convo).Association("Messages").Append(msg)
	return gormErr(err)
}

func gormResult(result *gorm.DB) error {
	return gormErr(result.Error)
}

func gormErr(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNotFound
	default:
		return ErrRepository
	}
}

package conversation

import (
	"context"

	"app/message"
	"app/user"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) Create(ctx context.Context, user1, user2 user.User) (Conversation, error) {
	return service.repo.Create(ctx, user1, user2)
}

func (service *Service) Find(ctx context.Context, id string) (Conversation, error) {
	return service.repo.Find(ctx, id)
}

func (service *Service) FindForUser(ctx context.Context, user user.User) ([]Conversation, error) {
	return service.repo.FindForUser(ctx, user)
}

func (service *Service) SendMessage(ctx context.Context, convo Conversation, author user.User, text string) (message.Message, error) {
	msg := message.Message{AuthorID: author.ID, Text: text}
	return msg, service.repo.AddMessage(ctx, convo, &msg)
}

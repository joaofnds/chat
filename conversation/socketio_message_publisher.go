package conversation

import (
	"app/message"
	"app/user"
	"encoding/json"

	"github.com/zishang520/socket.io/socket"
)

type SocketIOMessagePublisher struct {
	server *socket.Server
}

func NewSocketIOMessagePublisher(server *socket.Server) *SocketIOMessagePublisher {
	return &SocketIOMessagePublisher{server: server}
}

func (s *SocketIOMessagePublisher) PublishMessage(convo Conversation, author user.User, msg message.Message) error {
	p := Payload{Message: msg, Author: author}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	for _, user := range convo.Users {
		if user.ID == author.ID {
			continue
		}
		if err := s.server.To(socket.Room("user:"+user.ID)).Emit("msg", string(b)); err != nil {
			return err
		}
	}

	return nil
}

type Payload struct {
	Message message.Message `json:"message"`
	Author  user.User       `json:"author"`
}

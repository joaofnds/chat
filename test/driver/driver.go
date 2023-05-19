package driver

import (
	"app/conversation"
	"app/message"
	"app/user"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Driver struct {
	api *API
}

func NewDriver(url string) *Driver {
	return &Driver{api: NewAPI(url)}
}

func (d *Driver) CreateUser(name string) (user.User, error) {
	var u user.User

	return u, makeJSONRequest(func() (*http.Response, error) {
		return d.api.CreateUser(name)
	}, http.StatusCreated, &u)
}

func (d *Driver) GetUser(id string) (user.User, error) {
	var u user.User

	return u, makeJSONRequest(func() (*http.Response, error) {
		return d.api.GetUser(id)
	}, http.StatusOK, &u)
}

func (d *Driver) ListUsers() ([]user.User, error) {
	var users []user.User
	return users, makeJSONRequest(func() (*http.Response, error) {
		return d.api.ListUsers()
	}, http.StatusOK, &users)
}

func (d *Driver) GetConversation(id string) (conversation.Conversation, error) {
	var convo conversation.Conversation

	return convo, makeJSONRequest(func() (*http.Response, error) {
		return d.api.GetConversation(id)
	}, http.StatusOK, &convo)
}

func (d *Driver) CreateConversation(sender, receiver user.User) (conversation.Conversation, error) {
	var convo conversation.Conversation

	return convo, makeJSONRequest(func() (*http.Response, error) {
		return d.api.CreateConversation(sender.ID, receiver.ID)
	}, http.StatusCreated, &convo)
}

func (d *Driver) SendMessage(convo conversation.Conversation, author user.User, text string) (message.Message, error) {
	var msg message.Message

	return msg, makeJSONRequest(func() (*http.Response, error) {
		return d.api.SendMessage(convo.ID, author.ID, text)
	}, http.StatusCreated, &msg)
}

func makeJSONRequest(f func() (*http.Response, error), expectedStatus int, v any) error {
	res, err := f()
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != expectedStatus {
		return fmt.Errorf("expected status %d, got %d", expectedStatus, res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func makeRequest(f func() (*http.Response, error), expectedStatus int) error {
	res, err := f()
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != expectedStatus {
		return fmt.Errorf("expected status %d, got %d", expectedStatus, res.StatusCode)
	}

	return nil
}

package driver

import (
	"fmt"
	"net/http"
	"strings"

	"app/test/req"
)

type API struct {
	baseURL string
}

func NewAPI(baseURL string) *API {
	return &API{baseURL}
}

func (a API) CreateUser(name string) (*http.Response, error) {
	return req.Post(
		a.baseURL+"/users",
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(fmt.Sprintf(`{"name":%q}`, name)),
	)
}

func (a API) GetUser(id string) (*http.Response, error) {
	return req.Get(
		a.baseURL+"/users/"+id,
		map[string]string{"Accept": "application/json"},
	)
}

func (a API) ListUsers() (*http.Response, error) {
	return req.Get(
		a.baseURL+"/users",
		map[string]string{"Accept": "application/json"},
	)
}

func (a API) GetConversation(id string) (*http.Response, error) {
	return req.Get(
		a.baseURL+"/conversations/"+id,
		map[string]string{"Accept": "application/json"},
	)
}

func (a API) CreateConversation(senderID, receiverID string) (*http.Response, error) {
	return req.Post(
		a.baseURL+"/conversations",
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(fmt.Sprintf(`{"sender_id":%q,"receiver_id":%q}`, senderID, receiverID)),
	)
}

func (a API) SendMessage(conversationID, authorID, text string) (*http.Response, error) {
	return req.Post(
		a.baseURL+"/conversations/"+conversationID,
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(fmt.Sprintf(`{"author_id":%q,"text":%q}`, authorID, text)),
	)
}

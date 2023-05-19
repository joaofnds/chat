package main

import (
	"app/adapters/health"
	"app/adapters/http"
	"app/adapters/logger"
	"app/adapters/postgres"
	"app/adapters/websocket"

	"app/config"
	"app/conversation"
	"app/message"
	"app/user"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		logger.Module,
		health.Module,

		postgres.Module,
		http.Module,
		websocket.Module,

		user.Module,
		conversation.Module,
		message.Module,
	).Run()
}

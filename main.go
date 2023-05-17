package main

import (
	"app/adapters/health"
	"app/adapters/http"
	"app/adapters/logger"
	"app/adapters/metrics"
	"app/adapters/postgres"

	"app/config"
	"app/message"
	"app/user"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		logger.Module,
		metrics.Module,
		health.Module,

		http.Module,
		postgres.Module,

		user.Module,
		message.Module,
	).Run()
}

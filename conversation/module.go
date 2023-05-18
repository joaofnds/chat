package conversation

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"conversation",
	fx.Provide(NewService),

	fx.Provide(NewPostgresRepository),
	fx.Provide(func(repo *PostgresRepository) Repository { return repo }),

	fx.Provide(NewController),
	fx.Invoke(func(app *fiber.App, controller *Controller) { controller.Register(app) }),
)

package conversation

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"conversation",
	fx.Provide(NewService),

	fx.Provide(NewPostgresRepository),
	fx.Provide(func(repo *PostgresRepository) Repository { return repo }),
)

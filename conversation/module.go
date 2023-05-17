package conversation

import "go.uber.org/fx"

var Module = fx.Module(
	"conversation",
	fx.Invoke(AutoMigrate),
)

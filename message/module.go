package message

import "go.uber.org/fx"

var Module = fx.Module(
	"message",
	fx.Invoke(AutoMigrate),
)
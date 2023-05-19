package test

import (
	"github.com/zishang520/socket.io/socket"
	"go.uber.org/fx"
)

var TestSocketIO = fx.Provide(func() *socket.Server { return socket.NewServer(nil, nil) })

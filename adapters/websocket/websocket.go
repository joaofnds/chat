package websocket

import (
	"context"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/zishang520/socket.io/socket"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"websocket",
	fx.Provide(NewServer),
	fx.Invoke(HookServer),
	fx.Invoke(Setup),
	fx.Invoke(Register),
)

func NewServer() *socket.Server {
	return socket.NewServer(nil, nil)
}

func Register(app *fiber.App, server *socket.Server) {
	// app.Static("/", ".") // uncomment this to use test.html
	app.All("/socket.io", adaptor.HTTPHandler(server.ServeHandler(nil)))
}

func Setup(server *socket.Server, logger *zap.Logger) error {
	return server.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)

		err := client.On("user", func(data ...any) {
			id := data[0].(string)
			client.Join(socket.Room("user:" + id))
		})
		if err != nil {
			logger.Error("failed to listen to user event", zap.Error(err))
		}
	})
}

func HookServer(lifecycle fx.Lifecycle, server *socket.Server) {
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			server.Close(nil)
			return nil
		},
	})
}

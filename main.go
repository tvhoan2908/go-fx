package main

import (
	"context"
	"net/http"

	"github.com/tvhoan2908/go-fx/db"
	"github.com/tvhoan2908/go-fx/handler"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func newServer(lc fx.Lifecycle, userHandler *handler.UserHandler) *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		res := userHandler.GetUsers()
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":1323")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Close()
		},
	})

	return e
}

func main() {
	fx.New(
		fx.Provide(
			db.NewDb,
			handler.NewUserHandler,
		),
		fx.Invoke(newServer),
	).Run()
}

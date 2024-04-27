package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, ctx.QueryParam("id"))
	})

	e.Logger.Warn(e.Start(":3001"))
}
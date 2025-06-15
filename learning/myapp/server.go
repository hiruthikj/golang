package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"name"`
}

type UserDTO struct {
	Name    string
	Email   string
	IsAdmin bool
}

func main() {
	e := echo.New()

	e.GET("/healthcheck", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, ctx.QueryParam("id"))
	})

	e.POST("/users", func(c echo.Context) error {
		u := &User{}
		err := c.Bind(u)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		user := UserDTO{
			Name:    u.Name,
			Email:   u.Email,
			IsAdmin: false,
		}

		c.Logger().Infof("Created %v", user)

		return c.String(http.StatusCreated, "Created User")
	})

	if err := e.Start(":1323"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

package main

import (
	"net/http"

	"echo-crud/handler"
	"echo-crud/repositories"
	"echo-crud/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	r := repositories.New(e)
	u := usecase.New(r)
	handler.New(e, u)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}

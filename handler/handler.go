package handler

import (
	"echo-crud/handler/about"
	"echo-crud/handler/welcome"
	"echo-crud/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	U *usecase.Usecase
	E *echo.Echo
}

func New(e *echo.Echo, u *usecase.Usecase) {
	about.New(e, u)
	welcome.New(e, u)
}

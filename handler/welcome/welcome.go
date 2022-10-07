package welcome

import (
	"net/http"

	"echo-crud/usecase"

	"github.com/labstack/echo/v4"
)

type Welcome struct {
	u *usecase.Usecase
}

func New(e *echo.Echo, u *usecase.Usecase) {
	g := e.Group("/welcome")
	wfunc := Welcome{
		u: u,
	}

	g.GET("/1", wfunc.W1)
	g.GET("/2", wfunc.W2)
}

func (w *Welcome) W1(e echo.Context) error {
	return e.String(http.StatusOK, "Welcome 1")
}

func (w *Welcome) W2(e echo.Context) error {
	return e.String(http.StatusOK, "Welcome 2")
}

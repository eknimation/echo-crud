package about

import (
	"net/http"

	"echo-crud/usecase"

	"github.com/labstack/echo/v4"
)

type About struct {
	u *usecase.Usecase
}

func New(e *echo.Echo, u *usecase.Usecase) {
	g := e.Group("/about")
	afunc := About{
		u: u,
	}

	g.GET("/1", afunc.W1)
	g.GET("/2", afunc.W2)
}

func (w *About) W1(e echo.Context) error {
	result := w.u.U1()
	return e.String(http.StatusOK, result)
}

func (w *About) W2(e echo.Context) error {
	return e.String(http.StatusOK, "About 2")
}

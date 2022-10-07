package repositories

import "github.com/labstack/echo/v4"

type Repositiries struct {
	Echo *echo.Echo
}

func New(e *echo.Echo) *Repositiries {
	return &Repositiries{
		Echo: e,
	}
}

func (r *Repositiries) GetName(ID int) (string, error) {
	result := "Not found"
	if ID == 1 {
		result = "Jan"
	} else {
		result = "Doe"
	}

	return result, nil
}

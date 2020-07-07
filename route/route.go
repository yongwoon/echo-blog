package route

import (
	"github.com/labstack/echo/v4"
)

// Init Initialiize routes
func Init(e *echo.Echo) {

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.Validator = NewValidator()

	routeV1(e)
}

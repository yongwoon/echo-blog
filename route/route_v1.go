package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog-api/config"
)

func routeV1(e *echo.Echo) {
	v1 := e.Group(config.APIVersion1)

	v1.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

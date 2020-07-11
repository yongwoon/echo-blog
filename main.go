package main

import (
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
	"github.com/yongwoon/echo-blog/config/initializer"
	"github.com/yongwoon/echo-blog/db"
	"github.com/yongwoon/echo-blog/middleware"
	"github.com/yongwoon/echo-blog/route"
)

func main() {
	initializer.InitDotenv()

	db.Init()

	e := echo.New()

	e.Use(middleware.AccessLog())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))

	route.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}

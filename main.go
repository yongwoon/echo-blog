package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog-api/config/initializer"
	"github.com/yongwoon/echo-blog-api/route"
)

func main() {
	initializer.InitDotenv()

	e := echo.New()

	route.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}

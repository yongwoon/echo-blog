package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog/config/initializer"
	"github.com/yongwoon/echo-blog/db"
	"github.com/yongwoon/echo-blog/route"
)

func main() {
	initializer.InitDotenv()

	e := echo.New()

	db.Init()

	route.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}

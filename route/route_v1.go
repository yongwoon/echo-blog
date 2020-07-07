package route

import (
	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog/config"
	v1controller "github.com/yongwoon/echo-blog/controller/v1"
)

func routeV1(e *echo.Echo) {
	v1 := e.Group(config.APIVersion1)
	posts := v1.Group("/posts")

	postController := v1controller.NewPost()
	posts.GET("", postController.Index)
}

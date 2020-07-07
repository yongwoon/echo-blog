package route

import (
	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog/config"
	v1controller "github.com/yongwoon/echo-blog/controller/v1"
	"github.com/yongwoon/echo-blog/persistence"
)

func routeV1(e *echo.Echo) {
	v1 := e.Group(config.APIVersion1)
	posts := v1.Group("/posts")

	postPersistence := persistence.NewPostPersistence()
	postController := v1controller.NewPost(postPersistence)
	posts.GET("", postController.Index)
	posts.POST("", postController.Create)
}

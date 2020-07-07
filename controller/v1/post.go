package v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog-api/db"
	"github.com/yongwoon/echo-blog-api/model"
	"github.com/yongwoon/echo-blog-api/serializer"
)

type (
	// PostController postController
	PostController struct {
	}
)

// NewPost returns NewPost instance.
func NewPost() *PostController {
	return &PostController{}
}

// Index : api/v1/posts
func (p *PostController) Index(c echo.Context) error {
	db := db.DbManager()
	var posts []model.Post
	if err := db.Find(&posts); err != nil {
		fmt.Println(err)
	}

	serializer := serializer.PostsSerializer{posts}
	return c.JSON(http.StatusOK, serializer.Response())
}

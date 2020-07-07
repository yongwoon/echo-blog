package v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yongwoon/echo-blog/db"
	"github.com/yongwoon/echo-blog/form"
	"github.com/yongwoon/echo-blog/model"
	"github.com/yongwoon/echo-blog/serializer"
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

	return c.JSON(http.StatusOK, serializer.NewPostListResponse(posts))
}

// Create create post
func (p *PostController) Create(c echo.Context) error {
	db := db.DbManager()

	// ----------- params
	req, err := form.NewPost(c)
	if err != nil {
		fmt.Println("--- post create error ------")
	}

	// ----------- create
	post := model.Post{}
	post.Title = req.Post.Title
	post.Body = req.Post.Body

	db.Create(&post)

	return c.JSON(http.StatusCreated, serializer.NewPostResponse(&post))
}

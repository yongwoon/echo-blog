package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/yongwoon/echo-blog/form"
	"github.com/yongwoon/echo-blog/model"
	"github.com/yongwoon/echo-blog/repository"
	"github.com/yongwoon/echo-blog/serializer"
)

type (
	// PostController postController
	PostController struct {
		postRepository repository.PostRepository
	}
)

// NewPost returns NewPost instance.
func NewPost(pr repository.PostRepository) *PostController {
	return &PostController{
		postRepository: pr,
	}
}

// Index api/v1/posts
func (p *PostController) Index(c echo.Context) error {
	posts, err := p.postRepository.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, serializer.NewPostListResponse(posts))
}

// Show api/v1/posts
func (p *PostController) Show(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	post, err := p.postRepository.GetByID(postID)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, serializer.NewPostResponse(post))
}

// Create create post
func (p *PostController) Create(c echo.Context) error {
	req, err := form.NewPost(c)
	if err != nil {
		fmt.Println("--- post create error ------")
	}

	var post model.Post
	err = p.postRepository.Create(&post, req)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusCreated, serializer.NewPostResponse(&post))
}

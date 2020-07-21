package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/yongwoon/echo-blog/form"
	"github.com/yongwoon/echo-blog/model"
	"github.com/yongwoon/echo-blog/repository"
	"github.com/yongwoon/echo-blog/serializer"
	"github.com/yongwoon/echo-blog/utils"
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

// Index get all posts
// @Summary get posts
// @Description get all posts
// @Accept  json
// @Produce  json
// @Success 200 {object} serializer.PostListSerializer
// @Router /api/v1/posts [get]
func (p *PostController) Index(c echo.Context) error {
	posts, count, err := p.postRepository.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, serializer.NewPostListResponse(posts, count))
}

// Show api/v1/posts
// @Summary find post by id
// @Description find post by id
// @Accept  json
// @Produce  json
// @Param post_id path int true "Post ID"
// @Success 200 {object} serializer.SinglePostSerializer
// @Router /api/v1/posts/{post_id} [get]
func (p *PostController) Show(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	post, err := p.postRepository.GetByID(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if post == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, serializer.NewPostResponse(post))
}

// Create create post
// @Summary create new post
// @Description create new post
// @Accept  json
// @Produce  json
// @Param post body form.PostCreateReq true "Post Body"
// @Success 200 {object} serializer.SinglePostSerializer
// @Router /api/v1/posts/ [post]
func (p *PostController) Create(c echo.Context) error {
	req, err := form.NewPost(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	var post model.Post
	err = p.postRepository.Create(&post, req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, serializer.NewPostResponse(&post))
}

// Update post
// @Summary Update post
// @Description Update post
// @Accept  json
// @Produce  json
// @Param post_id path int true "Post ID"
// @Param post body form.PostUpdateReq true "Post Body"
// @Success 200 {object} serializer.SinglePostSerializer
// @Router /api/v1/posts/{post_id} [patch]
func (p *PostController) Update(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	post, err := p.postRepository.GetByID(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if post == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	req, err := form.UpdatePost(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err = p.postRepository.Update(post, req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, serializer.NewPostResponse(post))
}

// Delete post
// @Summary Delete post
// @Description Delete post
// @Accept  json
// @Produce  json
// @Param post_id path int true "Post ID"
// @Success 200 {object} serializer.SuccessSerializer
// @Router /api/v1/posts/{post_id} [delete]
func (p *PostController) Delete(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	post, err := p.postRepository.GetByID(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if post == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	err = p.postRepository.Delete(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, serializer.NewSuccessResponse())
}

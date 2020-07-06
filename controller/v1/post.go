package v1

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type (
	// IPost interface
	IPost interface {
		Index(echo.Context) error
	}

	// Post struct
	Post struct {
	}
)

// NewPost returns NewPost instance.
func NewPost() *Post {
	return &Post{}
}

// Index : api/vX/posts
func (p *Post) Index(c echo.Context) error {

	return c.JSON(http.StatusOK, "post all")
}

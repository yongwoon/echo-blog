package form

import (
	"github.com/labstack/echo/v4"
)

type (
	// PostCreateReq post create parameters
	PostCreateReq struct {
		Post struct {
			Title string `json:"title" validate:"required"`
			Body  string `json:"body" validate:"required"`
		} `json:"post"`
	}

	// PostUpdateReq post update parameters
	PostUpdateReq struct {
		Post struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"post"`
	}
)

// NewPost create post form
func NewPost(c echo.Context) (*PostCreateReq, error) {
	p := &PostCreateReq{}
	if err := c.Bind(p); err != nil {
		return nil, err
	}

	if err := c.Validate(p); err != nil {
		return nil, err
	}
	return p, nil
}

// UpdatePost create post form
func UpdatePost(c echo.Context) (*PostUpdateReq, error) {
	p := &PostUpdateReq{}
	if err := c.Bind(p); err != nil {
		return nil, err
	}

	if err := c.Validate(p); err != nil {
		return nil, err
	}
	return p, nil
}

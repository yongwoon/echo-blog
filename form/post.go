package form

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type (
	// PostCreateReq post create parameters
	PostCreateReq struct {
		Post struct {
			Title string `json:"title" form:"string" valid:"required, type(string)"`
			Body  string `json:"body" form:"string" valid:"required, type(string)"`
		} `json:"post" valid:"required"`
	}

	// PostUpdateReq post update parameters
	PostUpdateReq struct {
		Post struct {
			Title string `json:"title" form:"string" valid:"required, type(string)"`
			Body  string `json:"body" form:"string" valid:"required, type(string)"`
		} `json:"post" valid:"required"`
	}
)

// NewPost create post form
func NewPost(c echo.Context) (*PostCreateReq, error) {
	p := &PostCreateReq{}
	if err := c.Bind(p); err != nil {
		return nil, err
	}

	if !govalidator.MinStringLength(p.Post.Title, "1") {
		return nil, errors.New("title required")
	}

	if !govalidator.MinStringLength(p.Post.Body, "1") {
		return nil, errors.New("body required")
	}
	return p, nil
}

// UpdatePost create post form
// TODO: title, body が両方ない場合の  validation 追加
func UpdatePost(c echo.Context) (*PostUpdateReq, error) {
	p := &PostUpdateReq{}
	if err := c.Bind(p); err != nil {
		return nil, err
	}

	if !govalidator.MinStringLength(p.Post.Title, "1") {
		return nil, errors.New("title required")
	}

	if !govalidator.MinStringLength(p.Post.Body, "1") {
		return nil, errors.New("body required")
	}
	return p, nil
}

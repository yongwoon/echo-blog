package form

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type PostCreateReq struct {
	Post struct {
		Title string `json:"title" validate:"required"`
		Body  string `json:"body" validate:"required"`
	} `json:"post"`
}

// NewPost create post form
func NewPost(c echo.Context) (*PostCreateReq, error) {
	p := &PostCreateReq{}
	if err := c.Bind(p); err != nil {
		fmt.Println("------------- ERROR NEW POST ---------")
		return nil, err
	}

	if err := c.Validate(p); err != nil {
		fmt.Println("------------- ERROR POST Validation ---------")
		return nil, err
	}
	return p, nil
}

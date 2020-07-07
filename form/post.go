package form

import (
	"fmt"
	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PostCreateReq struct {
	Post struct {
		Title string `json:"title" validate:"required"`
		Body  string `json:"body" validate:"required"`
	} `json:"post"`
}

// NewCreditCard クレカのParameterの生成
func NewPost(c echo.Context) (*PostCreateReq, error) {
	p := &PostCreateReq{}
	if err := c.Bind(p); err != nil {
		fmt.Println("------------- ERROR NEW POST ---------")
		return nil, err
	}

	// if _, err := govalidator.ValidateStruct(f); err != nil {
	// 	return nil, err
	// }
	return p, nil
}

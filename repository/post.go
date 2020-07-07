package repository

import (
	"github.com/yongwoon/echo-blog/form"
	"github.com/yongwoon/echo-blog/model"
)

// PostRepository post repository interface
type PostRepository interface {
	GetAll() ([]model.Post, error)
	Create(*model.Post, *form.PostCreateReq) error
}

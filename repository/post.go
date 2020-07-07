package repository


import (
	"github.com/yongwoon/echo-blog/model"
)

type Store interface {
	All() (*model.Post, error)
	Create(*model.Post) error
}

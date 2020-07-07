package persistence

import (
	"fmt"

	"github.com/yongwoon/echo-blog/db"
	"github.com/yongwoon/echo-blog/form"
	"github.com/yongwoon/echo-blog/model"
	"github.com/yongwoon/echo-blog/repository"
)

// PostPersistence struct
type PostPersistence struct {
}

// NewPostPersistence post persistence
func NewPostPersistence() repository.PostRepository {
	return &PostPersistence{}
}

// GetAll get all posts
func (p PostPersistence) GetAll() ([]model.Post, error) {
	db := db.DbManager()

	var posts []model.Post
	if err := db.Find(&posts); err != nil {
		fmt.Println(err)
	}
	return posts, nil
}

// Create get all posts
func (p PostPersistence) Create(post *model.Post, req *form.PostCreateReq) error {
	db := db.DbManager()

	post.Title = req.Post.Title
	post.Body = req.Post.Body
	db.Create(&post)

	return nil
}

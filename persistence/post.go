package persistence

import (
	"github.com/jinzhu/gorm"
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
func (p PostPersistence) GetAll() ([]model.Post, int, error) {
	db := db.DbManager()

	var (
		posts []model.Post
		count int
	)

	if err := db.Find(&posts).Error; err != nil {
		return posts, 0, err
	}

	db.Model(&posts).Count(&count)

	return posts, count, nil
}

// GetByID get post by ID
func (p PostPersistence) GetByID(id int64) (*model.Post, error) {
	db := db.DbManager()

	var post model.Post

	err := db.First(&post, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

// Create get all posts
func (p PostPersistence) Create(post *model.Post, req *form.PostCreateReq) error {
	db := db.DbManager()

	post.Title = req.Post.Title
	post.Body = req.Post.Body

	tx := db.Begin()

	if err := tx.Create(&post).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update get all posts
func (p PostPersistence) Update(post *model.Post, req *form.PostUpdateReq) error {
	db := db.DbManager()

	post.Title = req.Post.Title
	post.Body = req.Post.Body

	tx := db.Begin()

	if err := tx.Model(post).Update(post).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Delete delete post
func (p PostPersistence) Delete(post *model.Post) error {
	db := db.DbManager()

	return db.Delete(post).Error
}

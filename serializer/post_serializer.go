package serializer

import (
	"time"

	"github.com/yongwoon/echo-blog/model"
)

type (
	PostSerializer struct {
		ID        uint      `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	SinglePostSerializer struct {
		Post *PostSerializer `json:"post"`
	}

	PostListSerializer struct {
		Posts      []*PostSerializer `json:"posts"`
		PostsCount int               `json:"postsCount"`
	}
)

// NewPostResponse return single post
func NewPostResponse(p *model.Post) *SinglePostSerializer {
	pr := new(PostSerializer)

	pr.ID = p.ID
	pr.Title = p.Title
	pr.Body = p.Body
	pr.CreatedAt = p.CreatedAt
	pr.UpdatedAt = p.UpdatedAt

	return &SinglePostSerializer{pr}
}

// NewPostListResponse return post list
func NewPostListResponse(posts []model.Post, count int) *PostListSerializer {
	r := new(PostListSerializer)
	r.Posts = make([]*PostSerializer, 0)
	for _, p := range posts {
		ar := new(PostSerializer)

		ar.ID = p.ID
		ar.Title = p.Title
		ar.Body = p.Body
		ar.CreatedAt = p.CreatedAt
		ar.UpdatedAt = p.UpdatedAt
		r.Posts = append(r.Posts, ar)
	}
	r.PostsCount = count

	return r
}

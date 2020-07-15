package serializer

import (
	"time"

	"github.com/yongwoon/echo-blog/model"
)

type (
	postSerializer struct {
		ID        uint      `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	// SinglePostSerializer single post serializer
	SinglePostSerializer struct {
		Post *postSerializer `json:"post"`
	}

	// PostListSerializer post list serializer
	PostListSerializer struct {
		Posts      []*postSerializer `json:"posts"`
		PostsCount int               `json:"postsCount"`
	}
)

// NewPostResponse return single post
func NewPostResponse(p *model.Post) *SinglePostSerializer {
	pr := new(postSerializer)

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
	r.Posts = make([]*postSerializer, 0)
	for _, p := range posts {
		ar := new(postSerializer)

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

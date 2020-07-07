package serializer

import (
	"time"

	"github.com/yongwoon/echo-blog/model"
)

type (
	postSerializer struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	singlePostSerializer struct {
		Post *postSerializer `json:"post"`
	}

	postListSerializer struct {
		Posts []*postSerializer `json:"posts"`
	}
)

// NewPostResponse return single post
func NewPostResponse(p *model.Post) *singlePostSerializer {
	pr := new(postSerializer)
	pr.ID = p.ID
	pr.Title = p.Title
	pr.Body = p.Body
	pr.CreatedAt = p.CreatedAt
	pr.UpdatedAt = p.UpdatedAt

	return &singlePostSerializer{pr}
}

// NewPostListResponse return post list
func NewPostListResponse(posts []model.Post) *postListSerializer {
	r := new(postListSerializer)
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

	return r
}

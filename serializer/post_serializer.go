package serializer

import (
	"github.com/yongwoon/echo-blog/model"
)

type (
	// PostSerializer one post record serializer
	PostSerializer struct {
		post model.Post
	}

	// PostsSerializer post list serializer
	PostsSerializer struct {
		Posts []model.Post
	}

	// PostResponse post response
	PostResponse struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Body      string `json:"body"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
)

// Response one post record response
func (s *PostSerializer) Response() PostResponse {
	return PostResponse{
		ID:        s.post.ID,
		Title:     s.post.Title,
		Body:      s.post.Body,
		CreatedAt: s.post.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: s.post.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
	}
}

// Response posts record response
func (s *PostsSerializer) Response() []PostResponse {
	response := []PostResponse{}
	for _, post := range s.Posts {
		serializer := PostSerializer{post}
		response = append(response, serializer.Response())
	}
	return response
}

package categoryparam

import (
	"encoding/json"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
)

type Response struct {
	category *entity.Category
}

func NewResponse(category *entity.Category) *Response {
	return &Response{category: category}
}
func (r *Response) GetCategory() *entity.Category {
	return r.category
}
func (r *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"category": r.GetCategory(),
	})
}

type ListResponse struct {
	categories []*entity.Category
}

func NewListResponse(categories []*entity.Category) *ListResponse {
	return &ListResponse{categories: categories}
}
func (lr *ListResponse) GetCategories() []*entity.Category {
	return lr.categories
}
func (lr *ListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"categories": lr.GetCategories(),
	})
}

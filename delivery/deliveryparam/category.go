package deliveryparam

import (
	"encoding/json"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"strings"
)

type CategoryRequest struct {
	title               string
	color               string
	authenticatedUserId uint
}

func NewCategoryRequest(title string, color string, authenticatedUserId uint) *CategoryRequest {
	return &CategoryRequest{
		title:               title,
		color:               color,
		authenticatedUserId: authenticatedUserId,
	}
}

func (c *CategoryRequest) GetTitle() string {
	return c.title
}
func (c *CategoryRequest) GetColor() string {
	return c.color
}
func (c *CategoryRequest) GetAuthenticatedUserId() uint {
	return c.authenticatedUserId
}
func (c *CategoryRequest) SetTitle(title string) {
	c.title = title
}
func (c *CategoryRequest) SetColor(color string) {
	c.color = color
}
func (c *CategoryRequest) SetAuthenticatedUserId(authenticatedUserId uint) {
	c.authenticatedUserId = authenticatedUserId
}
func (c *CategoryRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"title":               c.GetTitle(),
		"color":               c.GetColor(),
		"authenticatedUserId": c.GetAuthenticatedUserId(),
	})
}
func (c *CategoryRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		Title               string `json:"title"`
		Color               string `json:"color"`
		AuthenticatedUserId uint   `json:"authenticatedUserId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	c.SetTitle(aux.Title)
	c.SetColor(aux.Color)
	c.SetAuthenticatedUserId(aux.AuthenticatedUserId)

	return nil
}

type CategoryResponse struct {
	title      string
	categoryId uint
	error      error
}

func NewCategoryResponse(title string, categoryId uint, error error) *CategoryResponse {
	return &CategoryResponse{
		title:      title,
		categoryId: categoryId,
		error:      error,
	}
}
func (cr *CategoryResponse) GetTitle() string {
	return cr.title
}
func (cr *CategoryResponse) GetCategoryId() uint {
	return cr.categoryId
}
func (cr *CategoryResponse) GetError() error {
	return cr.error
}
func (cr *CategoryResponse) SetTitle(title string) {
	cr.title = title
}
func (cr *CategoryResponse) SetCategoryId(categoryId uint) {
	cr.categoryId = categoryId
}
func (cr *CategoryResponse) SetError(err error) {
	cr.error = err
}
func (cr *CategoryResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"title":      cr.GetTitle(),
		"categoryId": cr.GetCategoryId(),
		"error":      cr.GetError(),
	})
}
func (cr *CategoryResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		Title      string `json:"title"`
		CategoryId uint   `json:"categoryId"`
		Error      error  `json:"error"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	cr.SetTitle(aux.Title)
	cr.SetCategoryId(aux.CategoryId)
	cr.SetError(aux.Error)

	return nil
}

type CategoryListRequest struct {
	authenticatedUserId uint
}

func NewCategoryListRequest(authenticatedUserId uint) *CategoryListRequest {
	return &CategoryListRequest{authenticatedUserId: authenticatedUserId}
}
func (c *CategoryListRequest) GetAuthenticatedUserId() uint {
	return c.authenticatedUserId
}
func (c *CategoryListRequest) SetAuthenticatedUserId(authenticatedUserId uint) {
	c.authenticatedUserId = authenticatedUserId
}
func (c *CategoryListRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"authenticatedUserId": c.GetAuthenticatedUserId(),
	})
}
func (c *CategoryListRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		AuthenticatedUserId uint `json:"authenticatedUserId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	c.SetAuthenticatedUserId(aux.AuthenticatedUserId)

	return nil
}

type CategoryListResponse struct {
	categories []*entity.Category
}

func (c *CategoryListResponse) String() string {

	var categoriesStr strings.Builder = strings.Builder{}
	for _, cat := range c.GetCategories() {

		categoriesStr.WriteString(cat.String())
	}

	return categoriesStr.String()
}

func NewCategoryListResponse() *CategoryListResponse {
	return &CategoryListResponse{categories: make([]*entity.Category, 0)}
}
func (c *CategoryListResponse) GetCategories() []*entity.Category {
	return c.categories
}
func (c *CategoryListResponse) SetCategories(categories []*entity.Category) {
	c.categories = categories
}
func (c *CategoryListResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"categories": c.GetCategories(),
	})
}
func (c *CategoryListResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		Categories []*entity.Category `json:"categories"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	c.SetCategories(aux.Categories)

	return nil
}

package deliveryparam

import "encoding/json"

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

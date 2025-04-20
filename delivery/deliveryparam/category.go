package deliveryparam

import "encoding/json"

type Category struct {
	title string
	color string
}

type CategoryRequest struct {
	command  string
	category *Category
}

func (c *Category) GetTitle() string {
	return c.title
}
func (c *Category) GetColor() string {
	return c.color
}
func (c *Category) SetTitle(title string) {
	c.title = title
}
func (c *Category) SetColor(color string) {
	c.color = color
}
func (c *Category) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"title": c.GetTitle(),
		"color": c.GetColor(),
	})
}
func (c *Category) UnmarshalJSON(data []byte) error {
	var aux struct {
		Title string `json:"title"`
		Color string `json:"color"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	c.SetTitle(aux.Title)
	c.SetColor(aux.Color)

	return nil
}

func NewCategoryRequest(command string, title string, color string) *CategoryRequest {
	return &CategoryRequest{command: command, category: &Category{
		title: title,
		color: color,
	}}
}
func NewEmptyCategoryRequest() *CategoryRequest {
	return &CategoryRequest{
		command: "",
		category: &Category{
			title: "",
			color: "",
		},
	}
}
func (r *CategoryRequest) GetCommand() string {
	return r.command
}
func (r *CategoryRequest) SetCommand(command string) {
	r.command = command
}
func (r *CategoryRequest) GetCategory() *Category {
	return r.category
}
func (r *CategoryRequest) SetCategory(category *Category) {
	r.category = category
}
func (r *CategoryRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"command":  r.GetCommand(),
		"category": r.GetCategory(),
	})
}
func (r *CategoryRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		Command  string    `json:"command"`
		Category *Category `json:"category"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	r.SetCommand(aux.Command)
	r.SetCategory(aux.Category)

	return nil
}

package entity

type Category struct {
	id     uint
	title  string
	color  string
	userId uint
}

func NewCategory(id uint, title string, color string, userId uint) *Category {
	return &Category{
		id:     id,
		title:  title,
		color:  color,
		userId: userId,
	}
}

// GetId Getter method
func (c *Category) GetId() uint {
	return c.id
}

// GetTitle Getter method
func (c *Category) GetTitle() string {
	return c.title
}

// GetColor Getter method
func (c *Category) GetColor() string {
	return c.color
}

// GetUserId Getter method
func (c *Category) GetUserId() uint {
	return c.userId
}

// SetTitle Setter method
func (c *Category) SetTitle(title string) {
	c.title = title
}

// SetColor Setter method
func (c *Category) SetColor(color string) {
	c.color = color
}

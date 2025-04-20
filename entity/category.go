package entity

import (
	"encoding/json"
	"fmt"
)

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

func (c *Category) SetId(id uint) {
	c.id = id
}

// SetTitle Setter method
func (c *Category) SetTitle(title string) {
	c.title = title
}

// SetColor Setter method
func (c *Category) SetColor(color string) {
	c.color = color
}

func (c *Category) String() string {
	return fmt.Sprintf("Id: %d, Title: %s, Color: %s, UserId: %d\n",
		c.GetId(),
		c.GetTitle(),
		c.GetColor(),
		c.GetUserId(),
	)
}

func (c *Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"id":     c.GetId(),
		"title":  c.GetTitle(),
		"color":  c.GetColor(),
		"userId": c.GetUserId(),
	})
}

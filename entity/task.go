package entity

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	id         uint
	title      string
	dueDate    string
	categoryId uint
	isDone     bool
	userId     uint
}

func NewTask(id uint, title string, dueDate string, categoryId uint, userId uint) *Task {
	return &Task{
		id:         id,
		title:      title,
		dueDate:    dueDate,
		categoryId: categoryId,
		isDone:     false,
		userId:     userId,
	}
}

// GetId Getter method
func (t *Task) GetId() uint {
	return t.id
}

// GetTitle Getter method
func (t *Task) GetTitle() string {
	return t.title
}

// GetDueDate Getter method
func (t *Task) GetDueDate() string {
	return t.dueDate
}

// GetIsDone Getter method
func (t *Task) GetIsDone() bool {
	return t.isDone
}

// GetCategoryId Getter method
func (t *Task) GetCategoryId() uint {
	return t.categoryId
}

// GetUserId Getter method
func (t *Task) GetUserId() uint {
	return t.userId
}

// SetIsDone Setter method
func (t *Task) SetIsDone(isDone bool) {
	t.isDone = isDone
}

// SetDueDate Setter method
func (t *Task) SetDueDate(dueDate string) {
	t.dueDate = dueDate
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, DueDate: %s, IsDone: %v\n",
		t.id,
		t.title,
		t.dueDate,
		t.isDone,
	)
}

func (t *Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"id":         t.GetUserId(),
		"title":      t.GetTitle(),
		"dueDate":    t.GetDueDate(),
		"categoryId": t.GetCategoryId(),
		"isDone":     t.GetIsDone(),
		"userId":     t.GetUserId(),
	})
}

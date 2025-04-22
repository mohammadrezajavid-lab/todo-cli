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

// SetUserId Setter method
func (t *Task) SetUserId(userId uint) {
	t.userId = userId
}

func (t *Task) SetId(id uint) {
	t.id = id
}

func (t *Task) SetCategoryId(categoryId uint) {
	t.categoryId = categoryId
}

func (t *Task) SetTitle(title string) {
	t.title = title
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
		"id":         t.GetId(),
		"title":      t.GetTitle(),
		"dueDate":    t.GetDueDate(),
		"categoryId": t.GetCategoryId(),
		"isDone":     t.GetIsDone(),
		"userId":     t.GetUserId(),
	})
}

func (t *Task) UnmarshalJSON(data []byte) error {
	var aux struct {
		Id         uint   `json:"id"`
		Title      string `json:"title"`
		DueDate    string `json:"dueDate"`
		CategoryId uint   `json:"categoryId"`
		IsDone     bool   `json:"isDone"`
		UserId     uint   `json:"userId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetId(aux.Id)
	t.SetTitle(aux.Title)
	t.SetDueDate(aux.DueDate)
	t.SetCategoryId(aux.CategoryId)
	t.SetIsDone(aux.IsDone)
	t.SetUserId(aux.UserId)

	return nil
}

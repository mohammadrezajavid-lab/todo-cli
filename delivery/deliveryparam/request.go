package deliveryparam

import (
	"encoding/json"
)

type Task struct {
	title      string
	dueDate    string
	categoryId uint
}

func (t *Task) GetTitle() string {
	return t.title
}
func (t *Task) GetDueDate() string {
	return t.dueDate
}
func (t *Task) GetCategoryId() uint {
	return t.categoryId
}
func (t *Task) SetTitle(title string) {
	t.title = title
}
func (t *Task) SetDueDate(dueDate string) {
	t.dueDate = dueDate
}
func (t *Task) SetCategoryId(categoryId uint) {
	t.categoryId = categoryId
}
func (t *Task) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"title":      t.GetTitle(),
		"dueDate":    t.GetDueDate(),
		"categoryId": t.GetCategoryId(),
	})
}
func (t *Task) UnmarshalJSON(data []byte) error {
	var aux struct {
		Title      string `json:"title"`
		DueDate    string `json:"dueDate"`
		CategoryId uint   `json:"categoryId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetTitle(aux.Title)
	t.SetDueDate(aux.DueDate)
	t.SetCategoryId(aux.CategoryId)

	return nil
}

type Request struct {
	command string
	task    *Task
}

func NewRequest(command string, title string, dueDate string, categoryId uint) *Request {
	return &Request{command: command, task: &Task{
		title:      title,
		dueDate:    dueDate,
		categoryId: categoryId,
	}}
}

func NewEmptyRequest() *Request {
	return &Request{
		command: "",
		task: &Task{
			title:      "",
			dueDate:    "",
			categoryId: 0,
		},
	}
}

func (r *Request) GetCommand() string {
	return r.command
}
func (r *Request) SetCommand(command string) {
	r.command = command
}

func (r *Request) GetTask() *Task {
	return r.task
}
func (r *Request) SetTask(task *Task) {
	r.task = task
}

func (r *Request) UnmarshalJSON(data []byte) error {
	var aux struct {
		Command string `json:"command"`
		Task    *Task  `json:"task"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	r.SetCommand(aux.Command)
	r.SetTask(aux.Task)

	return nil
}

func (r *Request) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"command": r.GetCommand(),
		"task":    r.GetTask(),
	})
}

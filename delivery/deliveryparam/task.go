package deliveryparam

import "encoding/json"

type Task struct {
	title      string
	dueDate    string
	categoryId uint
}

type TaskRequest struct {
	command string
	task    *Task
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

func NewTaskRequest(command string, title string, dueDate string, categoryId uint) *TaskRequest {
	return &TaskRequest{command: command, task: &Task{
		title:      title,
		dueDate:    dueDate,
		categoryId: categoryId,
	}}
}
func NewEmptyTaskRequest() *TaskRequest {
	return &TaskRequest{
		command: "",
		task: &Task{
			title:      "",
			dueDate:    "",
			categoryId: 0,
		},
	}
}
func (r *TaskRequest) GetCommand() string {
	return r.command
}
func (r *TaskRequest) SetCommand(command string) {
	r.command = command
}
func (r *TaskRequest) GetTask() *Task {
	return r.task
}
func (r *TaskRequest) SetTask(task *Task) {
	r.task = task
}
func (r *TaskRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"command": r.GetCommand(),
		"task":    r.GetTask(),
	})
}
func (r *TaskRequest) UnmarshalJSON(data []byte) error {
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

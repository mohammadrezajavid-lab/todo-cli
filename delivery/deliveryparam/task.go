package deliveryparam

import "encoding/json"

type TaskRequest struct {
	title               string
	dueDate             string
	categoryId          uint
	authenticatedUserId uint
}

func NewTaskRequest(title string, dueDate string, categoryId uint, authenticatedUserId uint) *TaskRequest {
	return &TaskRequest{
		title:               title,
		dueDate:             dueDate,
		categoryId:          categoryId,
		authenticatedUserId: authenticatedUserId,
	}
}
func (t *TaskRequest) GetTitle() string {
	return t.title
}
func (t *TaskRequest) GetDueDate() string {
	return t.dueDate
}
func (t *TaskRequest) GetCategoryId() uint {
	return t.categoryId
}
func (t *TaskRequest) GetAuthenticatedUserId() uint {
	return t.authenticatedUserId
}
func (t *TaskRequest) SetTitle(title string) {
	t.title = title
}
func (t *TaskRequest) SetDueDate(dueDate string) {
	t.dueDate = dueDate
}
func (t *TaskRequest) SetCategoryId(categoryId uint) {
	t.categoryId = categoryId
}
func (t *TaskRequest) SetAuthenticatedUserId(authenticatedUserId uint) {
	t.authenticatedUserId = authenticatedUserId
}
func (t *TaskRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"title":               t.GetTitle(),
		"dueDate":             t.GetDueDate(),
		"categoryId":          t.GetCategoryId(),
		"authenticatedUserId": t.GetAuthenticatedUserId(),
	})
}
func (t *TaskRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		Title               string `json:"title"`
		DueDate             string `json:"dueDate"`
		CategoryId          uint   `json:"categoryId"`
		AuthenticatedUserId uint   `json:"authenticatedUserId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetTitle(aux.Title)
	t.SetDueDate(aux.DueDate)
	t.SetCategoryId(aux.CategoryId)
	t.SetAuthenticatedUserId(aux.AuthenticatedUserId)

	return nil
}

type TaskResponse struct {
	title  string
	taskId uint
	error  error
}

func NewTaskResponse(title string, taskId uint, err error) *TaskResponse {
	return &TaskResponse{
		title:  title,
		taskId: taskId,
		error:  err,
	}
}
func (t *TaskResponse) GetTitle() string {
	return t.title
}
func (t *TaskResponse) GetTaskId() uint {
	return t.taskId
}
func (t *TaskResponse) GetError() error {
	return t.error
}
func (t *TaskResponse) SetTitle(title string) {
	t.title = title
}
func (t *TaskResponse) SetTaskId(taskId uint) {
	t.taskId = taskId
}
func (t *TaskResponse) SetError(err error) {
	t.error = err
}
func (t *TaskResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"title":  t.GetTitle(),
		"taskId": t.GetTaskId(),
		"error":  t.GetError(),
	})
}
func (t *TaskResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		Title  string `json:"title"`
		TaskId uint   `json:"taskId"`
		Error  error  `json:"error"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetTitle(aux.Title)
	t.SetTaskId(aux.TaskId)
	t.SetError(aux.Error)

	return nil
}

type ListTaskRequest struct {
}

type ListTaskResponse struct {
}

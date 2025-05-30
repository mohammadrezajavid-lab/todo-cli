package deliveryparam

import (
	"encoding/json"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"strings"
)

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
	authenticatedUserId uint
}

func NewListTaskRequest(authenticatedUserId uint) *ListTaskRequest {
	return &ListTaskRequest{authenticatedUserId: authenticatedUserId}
}
func (t *ListTaskRequest) GetAuthenticatedUserId() uint {
	return t.authenticatedUserId
}
func (t *ListTaskRequest) SetAuthenticatedUserId(authenticatedUserId uint) {
	t.authenticatedUserId = authenticatedUserId
}
func (t *ListTaskRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"authenticatedUserId": t.GetAuthenticatedUserId(),
	})
}
func (t *ListTaskRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		AuthenticatedUserId uint `json:"authenticatedUserId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetAuthenticatedUserId(aux.AuthenticatedUserId)

	return nil
}

type ListTaskResponse struct {
	tasks []*entity.Task
}

func NewListTaskResponse() *ListTaskResponse {
	return &ListTaskResponse{
		tasks: make([]*entity.Task, 0),
	}
}
func (t *ListTaskResponse) GetTasks() []*entity.Task {
	return t.tasks
}
func (t *ListTaskResponse) SetTasks(tasks []*entity.Task) {
	t.tasks = tasks
}
func (t *ListTaskResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"tasks": t.GetTasks(),
	})
}
func (t *ListTaskResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		Tasks []*entity.Task `json:"tasks"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetTasks(aux.Tasks)

	return nil
}
func (t *ListTaskResponse) String() string {

	var categoriesStr strings.Builder = strings.Builder{}
	for _, cat := range t.GetTasks() {

		categoriesStr.WriteString(cat.String())
	}

	return categoriesStr.String()
}

type ListTaskByDateRequest struct {
	authenticatedUserId uint
	dueDate             string
}

func NewListTaskByDateRequest(authenticatedUserId uint, dueDate string) *ListTaskByDateRequest {
	return &ListTaskByDateRequest{
		authenticatedUserId: authenticatedUserId,
		dueDate:             dueDate,
	}
}
func (t *ListTaskByDateRequest) GetAuthenticatedUserId() uint {
	return t.authenticatedUserId
}
func (t *ListTaskByDateRequest) SetAuthenticatedUserId(authenticatedUserId uint) {
	t.authenticatedUserId = authenticatedUserId
}
func (t *ListTaskByDateRequest) GetDueDate() string {
	return t.dueDate
}
func (t *ListTaskByDateRequest) SetDueDate(dueDate string) {
	t.dueDate = dueDate
}
func (t *ListTaskByDateRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"authenticatedUserId": t.GetAuthenticatedUserId(),
		"dueDate":             t.GetDueDate(),
	})
}
func (t *ListTaskByDateRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		AuthenticatedUserId uint   `json:"authenticatedUserId"`
		DueDate             string `json:"dueDate"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetAuthenticatedUserId(aux.AuthenticatedUserId)
	t.SetDueDate(aux.DueDate)

	return nil
}

type ListTaskByDateResponse struct {
	tasks []*entity.Task
}

func NewListTaskByDateResponse() *ListTaskByDateResponse {
	return &ListTaskByDateResponse{
		tasks: make([]*entity.Task, 0),
	}
}
func (t *ListTaskByDateResponse) GetTasks() []*entity.Task {
	return t.tasks
}
func (t *ListTaskByDateResponse) SetTasks(tasks []*entity.Task) {
	t.tasks = tasks
}
func (t *ListTaskByDateResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"tasks": t.GetTasks(),
	})
}
func (t *ListTaskByDateResponse) UnmarshalJSON(data []byte) error {
	var aux struct {
		Tasks []*entity.Task `json:"tasks"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetTasks(aux.Tasks)

	return nil
}
func (t *ListTaskByDateResponse) String() string {

	var categoriesStr strings.Builder = strings.Builder{}
	for _, cat := range t.GetTasks() {

		categoriesStr.WriteString(cat.String())
	}

	return categoriesStr.String()
}

type ListTaskByStatusRequest struct {
	authenticatedUserId uint
	taskStatus          bool // true = Done & false = UnDone
}

func NewListTaskByStatusRequest(authenticatedUserId uint, taskStatus string) *ListTaskByStatusRequest {

	var status bool

	if taskStatus == "Done" {
		status = true
	} else {
		status = false
	}

	return &ListTaskByStatusRequest{authenticatedUserId: authenticatedUserId, taskStatus: status}
}
func (t *ListTaskByStatusRequest) GetAuthenticatedUserId() uint {
	return t.authenticatedUserId
}
func (t *ListTaskByStatusRequest) SetAuthenticatedUserId(authenticatedUserId uint) {
	t.authenticatedUserId = authenticatedUserId
}
func (t *ListTaskByStatusRequest) GetTaskStatus() bool {
	return t.taskStatus
}
func (t *ListTaskByStatusRequest) SetTaskStatus(taskStatus bool) {
	t.taskStatus = taskStatus
}
func (t *ListTaskByStatusRequest) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"authenticatedUserId": t.GetAuthenticatedUserId(),
		"taskStatus":          t.GetTaskStatus(),
	})
}
func (t *ListTaskByStatusRequest) UnmarshalJSON(data []byte) error {
	var aux struct {
		AuthenticatedUserId uint `json:"authenticatedUserId"`
		TaskStatus          bool `json:"taskStatus"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetAuthenticatedUserId(aux.AuthenticatedUserId)
	t.SetTaskStatus(aux.TaskStatus)

	return nil
}

type ListTaskByStatusResponse struct {
	tasks []*entity.Task
}

func NewListTaskByStatusResponse() *ListTaskByStatusResponse {

	return &ListTaskByStatusResponse{
		tasks: make([]*entity.Task, 0),
	}
}
func (t *ListTaskByStatusResponse) GetTasks() []*entity.Task {
	return t.tasks
}
func (t *ListTaskByStatusResponse) SetTasks(tasks []*entity.Task) {
	t.tasks = tasks
}
func (t *ListTaskByStatusResponse) MarshalJSON() ([]byte, error) {

	return json.Marshal(map[string]any{
		"tasks": t.GetTasks(),
	})
}
func (t *ListTaskByStatusResponse) UnmarshalJSON(data []byte) error {

	var aux struct {
		Tasks []*entity.Task `json:"tasks"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {

		return err
	}

	t.SetTasks(aux.Tasks)

	return nil
}
func (t *ListTaskByStatusResponse) String() string {

	var categoriesStr strings.Builder = strings.Builder{}
	for _, cat := range t.GetTasks() {

		categoriesStr.WriteString(cat.String())
	}

	return categoriesStr.String()
}

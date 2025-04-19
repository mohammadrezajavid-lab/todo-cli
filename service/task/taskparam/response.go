package taskparam

import (
	"encoding/json"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
)

type Response struct {
	task *entity.Task
}

func NewCreateTaskResponse(task *entity.Task) *Response {
	return &Response{task: task}
}
func (r *Response) GetTask() *entity.Task {
	return r.task
}

func (r *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"task": r.GetTask(),
	})
}

type ListResponse struct {
	tasks []*entity.Task
}

func NewListResponse(tasks []*entity.Task) *ListResponse {
	return &ListResponse{tasks: tasks}
}
func (lr *ListResponse) GetTasks() []*entity.Task {
	return lr.tasks
}
func (lr *ListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"tasks": lr.GetTasks(),
	})
}

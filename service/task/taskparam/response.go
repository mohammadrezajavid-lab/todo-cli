package taskparam

import (
	"encoding/json"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
)

type ResponseTask struct {
	task *entity.Task
}

func NewCreateTaskResponse(task *entity.Task) *ResponseTask {
	return &ResponseTask{task: task}
}
func (r *ResponseTask) GetTask() *entity.Task {
	return r.task
}
func (r *ResponseTask) MarshalJSON() ([]byte, error) {
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

type ListByDateResponse struct {
	tasks []*entity.Task
}

func NewListByDateResponse(tasks []*entity.Task) *ListByDateResponse {
	return &ListByDateResponse{
		tasks: tasks,
	}
}
func (lr *ListByDateResponse) GetTasks() []*entity.Task {
	return lr.tasks
}
func (lr *ListByDateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"tasks": lr.GetTasks(),
	})
}

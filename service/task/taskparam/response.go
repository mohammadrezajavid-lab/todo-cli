package taskparam

import "gocasts.ir/go-fundamentals/todo-cli/entity"

type Response struct {
	task *entity.Task
}

func NewCreateTaskResponse(task *entity.Task) *Response {
	return &Response{task: task}
}
func (r *Response) GetTask() *entity.Task {
	return r.task
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

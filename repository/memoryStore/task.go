package memoryStore

import (
	"gocasts.ir/go-fundamentals/todo-cli/entity"
)

type TaskMemory struct {
	tasks []*entity.Task
}

func NewTaskMemory() *TaskMemory {
	return &TaskMemory{make([]*entity.Task, 0)}
}

func (tm *TaskMemory) SetTasks(tasks []*entity.Task) {
	tm.tasks = tasks
}

func (tm *TaskMemory) GetTasks() []*entity.Task {
	return tm.tasks
}

func (tm *TaskMemory) CreateNewTask(t *entity.Task) (*entity.Task, error) {

	t.SetId(uint(len(tm.tasks) + 1))

	tm.SetTasks(append(tm.GetTasks(), t))

	return t, nil
}

func (tm *TaskMemory) ListUserTasks(userId uint) ([]*entity.Task, error) {

	userTasks := make([]*entity.Task, 0)
	for _, t := range tm.GetTasks() {

		if userId == t.GetUserId() {

			userTasks = append(userTasks, t)
		}
	}

	return userTasks, nil
}

func (tm *TaskMemory) ListTaskByDueDate(userId uint, dueDate string) ([]*entity.Task, error) {

	userTasks := make([]*entity.Task, 0)
	for _, t := range tm.GetTasks() {

		if userId == t.GetUserId() && dueDate == t.GetDueDate() {

			userTasks = append(userTasks, t)
		}
	}

	return userTasks, nil
}

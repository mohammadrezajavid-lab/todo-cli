package memoryStore

import (
	"gocasts.ir/go-fundamentals/todo-cli/constant"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/repository/filestore"
)

type TaskMemory struct {
	tasks     []*entity.Task
	taskStore *filestore.Store[entity.Task]
}

func NewTaskMemory() *TaskMemory {

	var taskMemory = &TaskMemory{tasks: make([]*entity.Task, 0), taskStore: new(filestore.Store[entity.Task])}

	var taskStore = filestore.NewStore[entity.Task](constant.TasksFile, constant.PermFile)
	taskMemory.SetTasks(append(taskMemory.GetTasks(), taskStore.Load(new(entity.Task))...))

	taskMemory.SetTaskStore(taskStore)

	return taskMemory
}

func (tm *TaskMemory) SetTasks(tasks []*entity.Task) {
	tm.tasks = tasks
}

func (tm *TaskMemory) GetTasks() []*entity.Task {
	return tm.tasks
}

func (tm *TaskMemory) GetTaskStore() *filestore.Store[entity.Task] {
	return tm.taskStore
}

func (tm *TaskMemory) SetTaskStore(taskStore *filestore.Store[entity.Task]) {
	tm.taskStore = taskStore
}

func (tm *TaskMemory) CreateNewTask(t *entity.Task) (*entity.Task, error) {

	// set id for new task
	t.SetId(uint(len(tm.tasks) + 1))

	// append new task to array of task
	tm.SetTasks(append(tm.GetTasks(), t))

	// write new task to database
	tm.GetTaskStore().Save(t)

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

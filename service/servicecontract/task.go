package servicecontract

import "gocasts.ir/go-fundamentals/todo-cli/entity"

type ServiceTaskRepository interface {
	CreateNewTask(t *entity.Task) (*entity.Task, error)
	ListUserTasks(userId uint) ([]*entity.Task, error)
}

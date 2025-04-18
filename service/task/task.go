package task

import (
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/service/task/taskparam"
)

type ServiceTaskRepository interface {
	CreateNewTask(t *entity.Task) (*entity.Task, error)
	ListUserTasks(userId uint) ([]*entity.Task, error)
}
type ServiceCategoryRepository interface {
	CheckCategoryId(categoryId uint, userId uint) (bool, error)
}
type Service struct {
	taskRepository     ServiceTaskRepository
	categoryRepository ServiceCategoryRepository
}

func NewService(taskRepository ServiceTaskRepository,
	categoryRepository ServiceCategoryRepository) *Service {
	return &Service{
		taskRepository:     taskRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *Service) CreateTask(taskReq *taskparam.Request) (*taskparam.Response, error) {

	//var ok, cErr = s.categoryRepository.CheckCategoryId(taskReq.Task.categoryId, taskReq.authenticatedUserId)
	//if cErr != nil {
	//	return nil, cErr
	//}
	//if !ok {
	//	return nil, fmt.Errorf("user does not have this categoryId: %d", taskReq.Task.categoryId)
	//}

	task, err := s.taskRepository.CreateNewTask(entity.NewTask(0, taskReq.GetTitle(), taskReq.GetDueDate(), taskReq.GetCategoryId(), taskReq.GetAuthenticatedUserId()))
	if err != nil {
		return nil, fmt.Errorf("can't create new task: %v", err)
	}

	return taskparam.NewCreateTaskResponse(task), nil
}

func (s *Service) ListTask(listReq *taskparam.ListRequest) (*taskparam.ListResponse, error) {

	tasks, err := s.taskRepository.ListUserTasks(listReq.GetUserId())

	if err != nil {

		return nil, fmt.Errorf("can't list tasks %v", err)
	}

	return taskparam.NewListResponse(tasks), nil
}

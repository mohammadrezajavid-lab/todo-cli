package task

import (
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/service/servicecontract"
	"gocasts.ir/go-fundamentals/todo-cli/service/task/taskparam"
)

type Service struct {
	taskRepository     servicecontract.ServiceTaskRepository
	categoryRepository servicecontract.ServiceCheckCategoryIdRepository
}

func NewService(taskRepository servicecontract.ServiceTaskRepository,
	categoryRepository servicecontract.ServiceCheckCategoryIdRepository) *Service {

	return &Service{
		taskRepository:     taskRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *Service) CreateTask(taskReq *taskparam.RequestTask) (*taskparam.ResponseTask, error) {

	var ok, cErr = s.categoryRepository.CheckCategoryId(taskReq.GetCategoryId(), taskReq.GetAuthenticatedUserId())
	if cErr != nil {
		return nil, cErr
	}
	if !ok {
		return nil, fmt.Errorf("user does not have this categoryId: %d", taskReq.GetCategoryId())
	}

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

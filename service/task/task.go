package task

import (
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
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

type Request struct {
	Task struct {
		title      string
		dueDate    string
		categoryId uint
	}

	authenticatedUserId uint
}

func (req *Request) GetTitle() string {
	return req.Task.title
}
func (req *Request) GetDueDate() string {
	return req.Task.dueDate
}
func (req *Request) GetCategoryId() uint {
	return req.Task.categoryId
}
func (req *Request) GetAuthenticatedUserId() uint {
	return req.authenticatedUserId
}

type Response struct {
	*entity.Task
}

func NewCreateTaskResponse(task *entity.Task) *Response {
	return &Response{task}
}

func (s *Service) CreateTask(taskReq *Request) (*Response, error) {

	var ok, cErr = s.categoryRepository.CheckCategoryId(taskReq.Task.categoryId, taskReq.authenticatedUserId)
	if cErr != nil {
		return nil, cErr
	}
	if !ok {
		return nil, fmt.Errorf("user does not have this categoryId: %d", taskReq.Task.categoryId)
	}

	task, err := s.taskRepository.CreateNewTask(entity.NewTask(0, taskReq.GetTitle(), taskReq.GetDueDate(), taskReq.GetCategoryId(), taskReq.GetAuthenticatedUserId()))
	if err != nil {
		return nil, fmt.Errorf("can't create new task: %v", err)
	}

	return NewCreateTaskResponse(task), nil
}

type ListRequest struct {
	userId uint
}

func NewListRequest(userId uint) *ListRequest {
	return &ListRequest{userId: userId}
}
func (lr *ListRequest) GetUserId() uint {
	return lr.userId
}

type ListResponse struct {
	tasks []*entity.Task
}

func NewListResponse(tasks []*entity.Task) *ListResponse {
	return &ListResponse{tasks: tasks}
}

func (s *Service) ListTask(listReq *ListRequest) (*ListResponse, error) {

	tasks, err := s.taskRepository.ListUserTasks(listReq.GetUserId())

	if err != nil {

		return nil, fmt.Errorf("can't list tasks %v", err)
	}

	return NewListResponse(tasks), nil
}

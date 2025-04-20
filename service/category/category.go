package category

import (
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/service/category/categoryparam"
	"gocasts.ir/go-fundamentals/todo-cli/service/servicecontract"
)

type Service struct {
	categoryRepository servicecontract.ServiceCategoryRepository
}

func NewService(categoryRepository servicecontract.ServiceCategoryRepository) *Service {
	return &Service{categoryRepository: categoryRepository}
}
func (s *Service) CreateCategory(categoryReq *categoryparam.Request) (*categoryparam.Response, error) {

	category, err := s.categoryRepository.CreateNewCategory(entity.NewCategory(0, categoryReq.GetTitle(), categoryReq.GetColor(), categoryReq.GetAuthenticatedUserId()))
	if err != nil {

		return nil, fmt.Errorf("can't create new category: %v", err)
	}

	return categoryparam.NewResponse(category), nil
}

func (s *Service) ListCategory(listReq *categoryparam.ListRequest) (*categoryparam.ListResponse, error) {

	categories, err := s.categoryRepository.ListUserCategories(listReq.GetUserId())
	if err != nil {

		return nil, fmt.Errorf("can't list categories %v", err)
	}

	return categoryparam.NewListResponse(categories), nil
}

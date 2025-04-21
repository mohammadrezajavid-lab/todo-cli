package user

import (
	"errors"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/pkg"
	"gocasts.ir/go-fundamentals/todo-cli/service/servicecontract"
	"gocasts.ir/go-fundamentals/todo-cli/service/user/userparam"
)

type Service struct {
	userRepository servicecontract.ServiceUserRepository
}

func NewService(userRepository servicecontract.ServiceUserRepository) *Service {
	return &Service{userRepository: userRepository}
}

func (s *Service) Login(reqUser *userparam.RequestUser) (*userparam.ResponseUser, error) {

	user, lErr := s.userRepository.Login(entity.NewUser(0, "", reqUser.GetEmail(), pkg.HashPassword(reqUser.GetPassword())))
	if lErr != nil {

		return nil, lErr
	}

	return userparam.NewResponseUser(user.GetId()), nil
}

func (s *Service) RegisterUser(reqUser *userparam.RequestRegisterUser) (*userparam.ResponseRegisterUser, error) {

	user, rErr := s.userRepository.CreateNewUser(entity.NewUser(0, reqUser.GetName(), reqUser.GetEmail(), reqUser.GetPassword()))
	if rErr != nil {

		return nil, errors.New("can't create new user")
	}

	return userparam.NewResponseRegisterUser(user.GetEmail()), nil
}

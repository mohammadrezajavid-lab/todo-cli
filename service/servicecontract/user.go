package servicecontract

import "gocasts.ir/go-fundamentals/todo-cli/entity"

type ServiceUserRepository interface {
	Login(user *entity.User) (*entity.User, error)
}

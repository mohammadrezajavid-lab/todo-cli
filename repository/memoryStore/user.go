package memoryStore

import (
	"errors"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
)

type UserMemory struct {
	users map[uint]*entity.User
}

func NewUserMemory() *UserMemory {
	return &UserMemory{users: make(map[uint]*entity.User)}
}

func (um *UserMemory) GetUsers() map[uint]*entity.User {
	return um.users
}

func (um *UserMemory) SetUsers(users map[uint]*entity.User) {
	um.users = users
}

func (um *UserMemory) CreateNewUser(u *entity.User) (*entity.User, error) {

	u.SetId(uint(len(um.GetUsers()) + 1))

	usersMap := um.GetUsers()
	usersMap[u.GetId()] = u
	um.SetUsers(usersMap)

	return u, nil
}

func (um *UserMemory) Login(user *entity.User) (*entity.User, error) {

	for _, u := range um.GetUsers() {
		if u.GetEmail() == user.GetEmail() && string(user.GetPassword()) == string(u.GetPassword()) {

			return u, nil
		}
	}

	return nil, errors.New("not found this user")
}

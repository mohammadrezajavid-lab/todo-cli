package memoryStore

import (
	"errors"
	"gocasts.ir/go-fundamentals/todo-cli/constant"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/repository/filestore"
)

type UserMemory struct {
	users     []*entity.User
	userStore *filestore.Store[entity.User]
}

func NewUserMemory() *UserMemory {

	var userMemory = &UserMemory{users: make([]*entity.User, 0), userStore: new(filestore.Store[entity.User])}

	var userStore = filestore.NewStore[entity.User](constant.UsersFile, constant.PermFile)
	userMemory.SetUsers(append(userMemory.GetUsers(), userStore.Load(new(entity.User))...))

	userMemory.SetUserStore(userStore)

	return userMemory
}

func (um *UserMemory) GetUsers() []*entity.User {
	return um.users
}

func (um *UserMemory) SetUsers(users []*entity.User) {
	um.users = users
}

func (um *UserMemory) GetUserStore() *filestore.Store[entity.User] {
	return um.userStore
}

func (um *UserMemory) SetUserStore(userStore *filestore.Store[entity.User]) {
	um.userStore = userStore
}

func (um *UserMemory) CreateNewUser(u *entity.User) (*entity.User, error) {

	// set id for new user
	u.SetId(uint(len(um.GetUsers()) + 1))

	// append new task to array of user
	um.SetUsers(append(um.GetUsers(), u))

	// write new user to database
	um.GetUserStore().Save(u)

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

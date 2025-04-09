package storage

import "gocasts.ir/go-fundamentals/todo-cli/entity"

type UserStore interface {
	Save(u *entity.User)
}

type TaskStore interface {
	Save(s *entity.Task)
}

type CategoryStore interface {
	Save(c *entity.Category)
}

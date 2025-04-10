package entity

import "fmt"

type Task struct {
	Title    string
	DueDate  string
	Category uint
	IsDone   bool
	UserId   uint
}

func (t *Task) String() string {
	return fmt.Sprintf("title: %s, userId: %d, dueDate: %s, isDone: %v, cat: %d\n",
		t.Title, t.UserId, t.DueDate, t.IsDone, t.Category)
}

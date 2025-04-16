package memoryStore

import (
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
)

type CategoryMemory struct {
	categories []*entity.Category
}

func NewCategoryMemory() *CategoryMemory {
	return &CategoryMemory{make([]*entity.Category, 0)}
}
func (cm *CategoryMemory) GetCategories() []*entity.Category {
	return cm.categories
}
func (cm *CategoryMemory) SetCategories(categories []*entity.Category) {
	cm.categories = categories
}

func (cm *CategoryMemory) CheckCategoryId(categoryId uint, userId uint) (bool, error) {
	for _, c := range cm.categories {
		if c.GetUserId() == userId && c.GetId() == categoryId {

			return true, nil
		}
	}

	return false, fmt.Errorf("this categoryId is not found: %d", categoryId)
}

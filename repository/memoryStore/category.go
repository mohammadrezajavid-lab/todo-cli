package memoryStore

import (
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/constant"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/repository/filestore"
)

type CategoryMemory struct {
	categories    []*entity.Category
	categoryStore *filestore.Store[entity.Category]
}

func NewCategoryMemory() *CategoryMemory {

	var categoryMemory = &CategoryMemory{categories: make([]*entity.Category, 0), categoryStore: new(filestore.Store[entity.Category])}

	var categoryStore *filestore.Store[entity.Category] = filestore.NewStore[entity.Category](constant.CategoriesFile, constant.PermFile)
	categoryMemory.SetCategories(append(categoryMemory.GetCategories(), categoryStore.Load(new(entity.Category))...))

	categoryMemory.SetCategoryStore(categoryStore)

	return categoryMemory
}

func (cm *CategoryMemory) GetCategories() []*entity.Category {
	return cm.categories
}

func (cm *CategoryMemory) SetCategories(categories []*entity.Category) {
	cm.categories = categories
}

func (cm *CategoryMemory) GetCategoryStore() *filestore.Store[entity.Category] {
	return cm.categoryStore
}

func (cm *CategoryMemory) SetCategoryStore(categoryStore *filestore.Store[entity.Category]) {
	cm.categoryStore = categoryStore
}

func (cm *CategoryMemory) CheckCategoryId(categoryId uint, userId uint) (bool, error) {
	for _, c := range cm.categories {
		if c.GetUserId() == userId && c.GetId() == categoryId {

			return true, nil
		}
	}

	return false, fmt.Errorf("this categoryId is not found: %d", categoryId)
}

func (cm *CategoryMemory) CreateNewCategory(c *entity.Category) (*entity.Category, error) {

	// set id for new category
	c.SetId(uint(len(cm.GetCategories()) + 1))

	// append new category to array of category
	cm.SetCategories(append(cm.GetCategories(), c))

	// write new category to database
	cm.GetCategoryStore().Save(c)

	return c, nil
}

func (cm *CategoryMemory) ListUserCategories(userId uint) ([]*entity.Category, error) {

	var categories = make([]*entity.Category, 0)

	for _, c := range cm.GetCategories() {
		if userId == c.GetUserId() {
			categories = append(categories, c)
		}
	}

	return categories, nil
}

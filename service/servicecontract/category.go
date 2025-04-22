package servicecontract

import "gocasts.ir/go-fundamentals/todo-cli/entity"

type ServiceCategoryRepository interface {
	CreateNewCategory(c *entity.Category) (*entity.Category, error)
	ListUserCategories(userId uint) ([]*entity.Category, error)
}

type ServiceCheckCategoryIdRepository interface {
	CheckCategoryId(categoryId uint, userId uint) (bool, error)
}

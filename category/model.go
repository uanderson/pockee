package category

import "github.com/uanderson/pockee/category/dao"

type CreateCategoryInput struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

type UpdateCategoryInput struct {
	ID   string `json:"id" validate:"required,min=20,max=20"`
	Name string `json:"name" validate:"required,min=2,max=255"`
}

type DeleteCategoryInput struct {
	ID string `json:"id" validate:"required,min=20,max=20"`
}

type CategoryOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToCategoryOutput(category dao.Category) CategoryOutput {
	return CategoryOutput{
		ID:   category.ID,
		Name: category.Name,
	}
}

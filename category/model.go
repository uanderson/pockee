package category

import (
	"github.com/uanderson/pockee/category/dao"
	"github.com/uanderson/pockee/model"
)

type CreateCategoryInput struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

type UpdateCategoryInput struct {
	CreateCategoryInput
	model.IdentifiableInput
}

type DeleteCategoryInput struct {
	model.IdentifiableInput
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

func ToCategoryOutputs(categories []dao.Category) []CategoryOutput {
	outputs := make([]CategoryOutput, len(categories))
	for i, c := range categories {
		outputs[i] = ToCategoryOutput(c)
	}

	return outputs
}

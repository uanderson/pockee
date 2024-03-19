// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dao

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, arg CreateCategoryParams) error
	DeleteCategory(ctx context.Context, arg DeleteCategoryParams) error
	ExistsCategoryByID(ctx context.Context, arg ExistsCategoryByIDParams) (bool, error)
	GetCategories(ctx context.Context, userID string) ([]Category, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error
}

var _ Querier = (*Queries)(nil)

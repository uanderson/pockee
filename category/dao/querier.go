// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dao

import (
	"context"
)

type Querier interface {
	GetCategories(ctx context.Context) ([]Category, error)
}

var _ Querier = (*Queries)(nil)

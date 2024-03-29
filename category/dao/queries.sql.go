// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package dao

import (
	"context"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name, parent_id, user_id) VALUES ($1, $2, $3, $4)
`

type CreateCategoryParams struct {
	ID       string
	Name     string
	ParentID *string
	UserID   string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.Exec(ctx, createCategory,
		arg.ID,
		arg.Name,
		arg.ParentID,
		arg.UserID,
	)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1 AND user_id = $2
`

type DeleteCategoryParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteCategory(ctx context.Context, arg DeleteCategoryParams) error {
	_, err := q.db.Exec(ctx, deleteCategory, arg.ID, arg.UserID)
	return err
}

const existsCategoryByID = `-- name: ExistsCategoryByID :one
SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1 AND user_id = $2)
`

type ExistsCategoryByIDParams struct {
	ID     string
	UserID string
}

func (q *Queries) ExistsCategoryByID(ctx context.Context, arg ExistsCategoryByIDParams) (bool, error) {
	row := q.db.QueryRow(ctx, existsCategoryByID, arg.ID, arg.UserID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getCategories = `-- name: GetCategories :many
SELECT id, name, parent_id, user_id FROM categories WHERE user_id = $1
`

func (q *Queries) GetCategories(ctx context.Context, userID string) ([]Category, error) {
	rows, err := q.db.Query(ctx, getCategories, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ParentID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories SET name = $1, parent_id = $2 WHERE id = $3 AND user_id = $4
`

type UpdateCategoryParams struct {
	Name     string
	ParentID *string
	ID       string
	UserID   string
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.Exec(ctx, updateCategory,
		arg.Name,
		arg.ParentID,
		arg.ID,
		arg.UserID,
	)
	return err
}

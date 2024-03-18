-- name: GetCategories :many
SELECT * FROM categories WHERE user_id = @user_id;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = @id AND user_id = @user_id;

-- name: ExistsCategoryByID :one
SELECT EXISTS(SELECT 1 FROM categories WHERE id = @id AND user_id = @user_id);

-- name: CreateCategory :exec
INSERT INTO categories (id, name, parent_id, user_id) VALUES (@id, @name, @parent_id, @user_id);

-- name: UpdateCategory :exec
UPDATE categories SET name = @name, parent_id = @parent_id WHERE id = @id AND user_id = @user_id;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = @id AND user_id = @user_id;

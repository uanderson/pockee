-- name: GetCategories :many
SELECT * FROM categories WHERE user_id = @user_id;

-- name: ExistsCategoryByID :one
SELECT EXISTS(SELECT 1 FROM categories WHERE id = @id AND user_id = @user_id);

-- name: CreateCategory :exec
INSERT INTO categories (id, name, user_id) VALUES (@id, @name, @user_id);

-- name: UpdateCategory :exec
UPDATE categories SET name = @name WHERE id = @id AND user_id = @user_id;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = @id AND user_id = @user_id;

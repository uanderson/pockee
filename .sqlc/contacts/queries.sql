-- name: GetContacts :many
SELECT * FROM contacts WHERE user_id = @user_id AND deleted_at IS NULL ORDER BY name;

-- name: GetContactByID :one
SELECT * FROM contacts WHERE id = @id AND user_id = @user_id AND deleted_at IS NULL;

-- name: ExistsContactByID :one
SELECT EXISTS(SELECT 1 FROM contacts WHERE id = @id AND user_id = @user_id AND deleted_at IS NULL);

-- name: CreateContact :exec
INSERT INTO contacts (id, name, email, phone, pix_key, user_id)
VALUES (@id, @name, @email, @phone, @pix_key, @user_id);

-- name: CreateContactHistory :exec
INSERT INTO contact_histories (id, name, email, phone, pix_key, effective_at, contact_id)
VALUES (@id, @name, @email, @phone, @pix_key, @effective_at, @contact_id);

-- name: UpdateContact :exec
UPDATE contacts SET name = @name, email = @email, phone = @phone, pix_key = @pix_key
WHERE id = @id AND user_id = @user_id;

-- name: SoftDeleteContact :exec
UPDATE contacts SET deleted_at = @deleted_at WHERE id = @id AND user_id = @user_id;
-- name: DeleteSettingByKey :exec
DELETE FROM settings WHERE key = @key AND user_id = @user_id;

-- name: ExistsSettingByKey :one
SELECT EXISTS (SELECT 1 FROM settings WHERE key = @key AND user_id = @user_id);

-- name: GetSettingByKey :one
SELECT * FROM settings WHERE key = @key AND user_id = @user_id;

-- name: UpdateSetting :exec
INSERT INTO settings (id, key, value, user_id) VALUES (@id, @key, @value, @user_id)
ON CONFLICT (key, user_id) DO UPDATE SET value = @value;

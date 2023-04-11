-- name: GetSettingByKey :one
SELECT * FROM settings WHERE key = @key;

-- name: GetUserSettingByKey :one
SELECT * FROM user_settings WHERE key = @key AND user_id = @user_id;

-- name: UpdateUserSetting :exec
INSERT INTO user_settings (id, key, value, user_id)
VALUES (@id, @key, @value, @user_id) ON CONFLICT (key, user_id) DO
UPDATE SET value = @value;

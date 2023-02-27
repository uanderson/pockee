-- name: GetSettingByKey :one
SELECT * FROM settings WHERE key = @key;

-- name: UpdateSetting :exec
INSERT INTO settings (id, key, value)
VALUES (@id, @key, @value) ON CONFLICT (key) DO
UPDATE SET value = @value;

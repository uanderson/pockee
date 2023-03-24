CREATE TABLE IF NOT EXISTS user_settings
(
    id      CHAR(20)     NOT NULL PRIMARY KEY,
    key     VARCHAR(255) NOT NULL,
    value   TEXT         NOT NULL,
    user_id CHAR(20)     NOT NULL,
    CONSTRAINT settings_key_user_id_uq UNIQUE (key, user_id)
);

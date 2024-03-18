CREATE TABLE IF NOT EXISTS settings
(
    id    CHAR(20)     NOT NULL PRIMARY KEY,
    key   VARCHAR(255) NOT NULL,
    value TEXT         NOT NULL,
    CONSTRAINT settings_key_uq UNIQUE (key)
);

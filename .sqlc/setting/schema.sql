CREATE TABLE settings
(
    id    CHAR(20)     NOT NULL PRIMARY KEY,
    key   VARCHAR(255) NOT NULL,
    value TEXT         NOT NULL
);

CREATE TABLE user_settings
(
    id      CHAR(20)     NOT NULL PRIMARY KEY,
    key     VARCHAR(255) NOT NULL,
    value   TEXT         NOT NULL,
    user_id CHAR(20)     NOT NULL
);
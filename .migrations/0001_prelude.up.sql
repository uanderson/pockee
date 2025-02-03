CREATE TABLE IF NOT EXISTS categories
(
  id      CHAR(20)     NOT NULL PRIMARY KEY,
  name    VARCHAR(255) NOT NULL,
  user_id CHAR(28)     NOT NULL
);

CREATE TABLE IF NOT EXISTS settings
(
  id      CHAR(20)     NOT NULL PRIMARY KEY,
  key     VARCHAR(255) NOT NULL,
  value   TEXT         NOT NULL,
  user_id CHAR(28)     NOT NULL,
  UNIQUE (key, user_id)
);
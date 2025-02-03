CREATE TABLE IF NOT EXISTS categories
(
  id      CHAR(20)     NOT NULL PRIMARY KEY,
  name    VARCHAR(255) NOT NULL,
  user_id CHAR(28)     NOT NULL
);
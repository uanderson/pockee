CREATE TABLE categories
(
  id        CHAR(20)     NOT NULL PRIMARY KEY,
  name      VARCHAR(255) NOT NULL,
  parent_id CHAR(20),
  user_id   CHAR(28)     NOT NULL
);
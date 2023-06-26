CREATE TABLE banks
(
  id   CHAR(3)      NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  ispb CHAR(8)      NOT NULL
);

CREATE TABLE accounts
(
  id      CHAR(20)     NOT NULL PRIMARY KEY,
  name    VARCHAR(255) NOT NULL,
  bank_id CHAR(20)     NOT NULL,
  user_id CHAR(28)     NOT NULL
);
CREATE TABLE IF NOT EXISTS contacts
(
  id         VARCHAR(20)  NOT NULL PRIMARY KEY,
  name       VARCHAR(100) NOT NULL,
  email      VARCHAR(100),
  phone      VARCHAR(20),
  pix_key    VARCHAR(255),
  deleted_at TIMESTAMP,
  user_id    VARCHAR(28)  NOT NULL
);

CREATE TABLE IF NOT EXISTS contact_histories
(
  id           VARCHAR(20)  NOT NULL PRIMARY KEY,
  name         VARCHAR(100) NOT NULL,
  email        VARCHAR(100),
  phone        VARCHAR(20),
  pix_key      VARCHAR(255),
  effective_at TIMESTAMP    NOT NULL,
  contact_id   VARCHAR(20)  NOT NULL REFERENCES contacts (id)
);
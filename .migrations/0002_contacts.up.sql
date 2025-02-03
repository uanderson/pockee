CREATE TABLE IF NOT EXISTS contacts
(
  id         VARCHAR(20)  NOT NULL PRIMARY KEY,
  name       VARCHAR(255) NOT NULL,
  email      VARCHAR(255) NULL,
  phone      VARCHAR(13)  NULL,
  pix_key    VARCHAR(255) NULL,
  deleted_at TIMESTAMP    NULL,
  user_id    VARCHAR(28)  NOT NULL
);

CREATE TABLE IF NOT EXISTS contact_histories
(
  id           VARCHAR(20)  NOT NULL PRIMARY KEY,
  name         VARCHAR(255) NOT NULL,
  email        VARCHAR(255) NULL,
  phone        VARCHAR(20)  NULL,
  pix_key      VARCHAR(255) NULL,
  effective_at TIMESTAMP    NOT NULL,
  contact_id   VARCHAR(20)  NOT NULL REFERENCES contacts (id)
);
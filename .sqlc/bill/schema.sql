CREATE TABLE recurring_bills
(
  id          CHAR(20)       NOT NULL PRIMARY KEY,
  description TEXT           NOT NULL,
  amount      NUMERIC(12, 2) NOT NULL,
  frequency   VARCHAR(20)    NOT NULL, -- DAILY, MONTHLY, YEARLY
  start_at    DATE           NOT NULL,
  end_at      DATE           NOT NULL,
  contact_id  CHAR(20)       NOT NULL,
  category_id CHAR(20)       NOT NULL,
  user_id     CHAR(28)       NOT NULL
);

CREATE TABLE bills
(
  id                CHAR(20)       NOT NULL PRIMARY KEY,
  description       TEXT           NOT NULL,
  amount            DECIMAL(12, 2) NOT NULL,
  status            VARCHAR(20)    NOT NULL, -- PENDING, PAID
  due_at            TIMESTAMP      NOT NULL,
  paid_at           TIMESTAMP,
  category_id       CHAR(20)       NOT NULL,
  contact_id        CHAR(20)       NOT NULL,
  user_id           CHAR(28)       NOT NULL
);

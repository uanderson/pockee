CREATE TABLE IF NOT EXISTS bills
(
  id          VARCHAR(20)      NOT NULL PRIMARY KEY,
  description TEXT             NOT NULL,
  type        VARCHAR(6)       NOT NULL,
  due_at      DATE             NOT NULL,
  amount      DOUBLE PRECISION NOT NULL,
  status      VARCHAR(20)      NOT NULL,
  category_id VARCHAR(20)      NOT NULL REFERENCES categories (id),
  contact_id  VARCHAR(20)      NOT NULL REFERENCES contacts (id),
  user_id     VARCHAR(28)      NOT NULL
);

CREATE TABLE IF NOT EXISTS recurring_bills
(
  id          VARCHAR(20)      NOT NULL PRIMARY KEY,
  description TEXT             NOT NULL,
  type        VARCHAR(6)       NOT NULL,
  start_at    DATE             NOT NULL,
  end_at      DATE             NULL,
  amount      DOUBLE PRECISION NOT NULL,
  interval    VARCHAR(20)      NOT NULL,
  contact_id  VARCHAR(20)      NOT NULL REFERENCES contacts (id),
  category_id VARCHAR(20)      NOT NULL REFERENCES categories (id),
  user_id     VARCHAR(28)      NOT NULL
);

CREATE TABLE IF NOT EXISTS boletos
(
  id      VARCHAR(20) NOT NULL PRIMARY KEY,
  barcode VARCHAR(44) NOT NULL,
  bill_id VARCHAR(20) NOT NULL REFERENCES bills (id)
);

CREATE TABLE IF NOT EXISTS payments
(
  id      VARCHAR(20)      NOT NULL PRIMARY KEY,
  date    DATE             NOT NULL,
  amount  DOUBLE PRECISION NOT NULL,
  bill_id VARCHAR(20)      NOT NULL REFERENCES bills (id)
);

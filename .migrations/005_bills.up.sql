-- Increase the length of the "user_id" column
ALTER TABLE user_settings ALTER COLUMN user_id TYPE CHAR(28);

-- Create account related tables
CREATE TABLE categories
(
  id        CHAR(20)     NOT NULL PRIMARY KEY,
  name      VARCHAR(255) NOT NULL,
  parent_id CHAR(20),
  user_id   CHAR(28),
  CONSTRAINT categories_categories_fkey FOREIGN KEY (parent_id) REFERENCES categories (id)
);

-- Create account and bank tables
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
  user_id CHAR(28)     NOT NULL,
  CONSTRAINT accounts_banks_fkey FOREIGN KEY (bank_id) REFERENCES banks (id)
);

-- Create user's contacts tables
CREATE TABLE contacts
(
  id      CHAR(20)     NOT NULL PRIMARY KEY,
  name    VARCHAR(100) NOT NULL,
  email   VARCHAR(100) NOT NULL,
  phone   VARCHAR(20)  NOT NULL,
  tax_id  VARCHAR(14)  NOT NULL,
  user_id CHAR(28)     NOT NULL
);

CREATE TABLE bank_details
(
  id             CHAR(20)     NOT NULL PRIMARY KEY,
  branch_number  INT          NOT NULL,
  account_number INT          NOT NULL,
  account_type   CHAR(1)      NOT NULL,
  pix_key        VARCHAR(255) NOT NULL,
  main           BOOLEAN      NOT NULL,
  bank_id        CHAR(3)      NOT NULL,
  contact_id     CHAR(20)     NOT NULL,
  CONSTRAINT bank_details_banks_fkey FOREIGN KEY (contact_id) REFERENCES banks (id),
  CONSTRAINT bank_details_contacts_fkey FOREIGN KEY (contact_id) REFERENCES contacts (id) ON DELETE CASCADE
);

-- Create bill' tables
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
  user_id     CHAR(28)       NOT NULL,
  CONSTRAINT recurring_bills_contacts_fkey FOREIGN KEY (contact_id) REFERENCES contacts (id),
  CONSTRAINT recurring_bills_categories_fkey FOREIGN KEY (category_id) REFERENCES categories (id)
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
  user_id           CHAR(28)       NOT NULL,
  CONSTRAINT bills_contacts_fkey FOREIGN KEY (contact_id) REFERENCES contacts (id),
  CONSTRAINT bills_categories_fkey FOREIGN KEY (category_id) REFERENCES categories (id)
);

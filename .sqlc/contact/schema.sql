CREATE TABLE contacts
(
    id      CHAR(20)     NOT NULL PRIMARY KEY,
    name    VARCHAR(100) NOT NULL,
    email   VARCHAR(100) NOT NULL,
    phone   VARCHAR(20)  NOT NULL,
    tax_id  VARCHAR(14)  NOT NULL,
    user_id CHAR(20)     NOT NULL
);

CREATE TABLE banks
(
    id   CHAR(3)      NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ispb CHAR(8)      NOT NULL
);

CREATE TABLE bank_details
(
    id             CHAR(20)     NOT NULL PRIMARY KEY,
    branch_number  INT          NOT NULL,
    account_number INT          NOT NULL,
    account_type   CHAR(1)      NOT NULL,
    pix_key        VARCHAR(255) NOT NULL,
    is_primary     BIT          NOT NULL,
    bank_id        CHAR(3)      NOT NULL,
    contact_id     CHAR(20)     NOT NULL
);

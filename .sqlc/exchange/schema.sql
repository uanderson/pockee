CREATE TABLE exchange_rates
(
    id         CHAR(20)         NOT NULL PRIMARY KEY,
    date       DATE             NOT NULL,
    source     CHAR(3)          NOT NULL,
    target     CHAR(3)          NOT NULL,
    rate       DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP        NOT NULL
);

CREATE TABLE exchange_currencies
(
    id     CHAR(20) NOT NULL PRIMARY KEY,
    source CHAR(3)  NOT NULL,
    target CHAR(3)  NOT NULL
);
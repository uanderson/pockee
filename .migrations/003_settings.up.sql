CREATE TABLE IF NOT EXISTS settings
(
    id    CHAR(20)     NOT NULL PRIMARY KEY,
    key   VARCHAR(255) NOT NULL,
    value TEXT         NOT NULL,
    CONSTRAINT settings_key_uq UNIQUE (key)
);

-- Inserts some initial settings
INSERT INTO settings (id, key, value)
VALUES ('Jx60PQTZxVy7yfntmCdA', 'exchange.cron', '0 9-17 * * 1-5'),
       ('4zPC6D7FDjcvQ8coTFqP', 'pocketsmith.cron', '10 9-17 * * 1-5');

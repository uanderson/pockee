-- Drops the unique constraint for date, source and target
ALTER TABLE exchange_rates DROP CONSTRAINT IF EXISTS date_source_target_uq;

/*
 * Adds the created at column for a more precise collection
 * time. It's important to notice that date is no longer
 * relevant for this table, but is kept not to break stuff.
 * Date will be scheduled for removal at a later time.
 */
 ALTER TABLE exchange_rates ADD COLUMN created_at TIMESTAMP;

-- Updates the created at column with the original date of the conversion
UPDATE exchange_rates SET created_at = date;

-- Makes the created at column not null
ALTER TABLE exchange_rates ALTER COLUMN created_at SET NOT NULL;

-- Creates a new table to map currencies that need to be converted
CREATE TABLE exchange_currencies
(
    id     CHAR(20) NOT NULL PRIMARY KEY,
    source CHAR(3)  NOT NULL,
    target CHAR(3)  NOT NULL
);

-- Inserts some initial currencies
INSERT INTO exchange_currencies (id, source, target)
VALUES ('74QlJMSeyNPCpEd9Lxni', 'CAD', 'BRL'),
       ('BQqIIVHyL9y3DqWtJCJy', 'EUR', 'BRL'),
       ('5Zx8f87i2poUJq0Tj9ay', 'GBP', 'BRL'),
       ('CG9aABakbxA03qnFZw8G', 'USD', 'BRL');

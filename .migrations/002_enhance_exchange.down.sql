-- Resets the exchange rates table to the previous state
ALTER TABLE exchange_rates DROP COLUMN created_at;
ALTER TABLE exchange_rates ADD CONSTRAINT date_source_target_uq UNIQUE (date, source, target);

-- Drop the exchange currencies table
DROP TABLE exchange_currencies;
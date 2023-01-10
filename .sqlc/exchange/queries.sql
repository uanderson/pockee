-- name: GetExchangeRate :one
SELECT *
FROM exchange_rates
WHERE date = now():: date;

-- name: GetExchangeRateForConversion :one
SELECT *
FROM exchange_rates
WHERE date = @date
  AND source = @source
  AND target = @target;

-- name: UpdateExchangeRate :exec
INSERT
INTO exchange_rates (id, date, source, target, rate)
VALUES (@id, @date, @source, @target, @rate)
ON CONFLICT (date, source, target) DO
UPDATE SET rate = @rate;

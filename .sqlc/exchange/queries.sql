-- name: GetExchangeRateForConversion :one
SELECT *
  FROM exchange_rates
 WHERE created_at::date = @created_at::date
   AND source = @source
   AND target = @target
 ORDER BY created_at DESC
 LIMIT 1;

-- name: CreateExchangeRate :exec
INSERT INTO exchange_rates (id, date, source, target, rate, created_at)
VALUES (@id, @date, @source, @target, @rate, @created_at);

-- name: GetExchangeCurrencies :many
SELECT * FROM exchange_currencies;
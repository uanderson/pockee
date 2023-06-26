-- name: GetRecurringBillsByRange :many
SELECT * FROM recurring_bills WHERE start_at::date = @date AND start_at::date = @date AND end_at::date <= @date;

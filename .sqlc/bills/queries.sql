-- name: CreateBill :exec
INSERT INTO bills (id, description, type, due_at, amount, status, category_id, contact_id, user_id)
VALUES (@id, @description, @type, @due_at, @amount, @status, @category_id, @contact_id, @user_id);


-- name: CreateRecurringBill :exec
INSERT INTO recurring_bills (id, description, type, start_at, end_at, amount, interval, contact_id, category_id, user_id)
VALUES (@id, @description, @type, @start_at, @end_at, @amount, @interval, @contact_id, @category_id, @user_id);

-- name: CreateBoleto :exec
INSERT INTO boletos (id, barcode, bill_id) VALUES (@id, @barcode, @bill_id);
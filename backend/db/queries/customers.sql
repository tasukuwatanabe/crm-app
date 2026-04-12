-- name: GetCustomers :many
SELECT * FROM customers
ORDER BY created_at DESC;

-- name: GetCustomersByID :one
SELECT * FROM customers
WHERE id = $1;

-- name: CreateCustomer :one
INSERT INTO customers (name, email, phone)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
SET name = $1, email = $2, phone = $3, updated_at = NOW()
WHERE id = $4
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;

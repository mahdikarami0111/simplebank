-- name: CreateAccount :one
INSERT INTO account (owner, balance, currency) VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountUpdate :one
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccount :many
SELECT * FROM account 
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: AddAcountBalance :one
UPDATE account
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account WHERE id = $1;
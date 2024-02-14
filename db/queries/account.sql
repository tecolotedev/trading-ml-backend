-- name: CreateAccount :one
INSERT INTO accounts (
    balance,
    currency,
    user_id
)
VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
where id =$1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
where id =$1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
WHERE user_id = $1
ORDER BY id 
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts
SET balance =$2
WHERE id = $1
RETURNING *;
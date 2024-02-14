-- name: CreateTransfer :one
INSERT INTO transfers (
  amount,
  reason,
  created_at,
  account_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    account_id = $1
ORDER BY id;
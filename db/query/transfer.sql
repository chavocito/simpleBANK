-- name: CreateTransfers :one
INSERT INTO transfers (
    id, from_account_id, to_account_id, amount
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY from_account_id;

-- name: UpdateTransfers :exec
UPDATE transfers
set from_account_id = $2,
    to_account_id = $3,
    amount = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;
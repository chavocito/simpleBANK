-- name: CreateEntries :one
INSERT INTO entries (
    id, account_id, amount
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY account_id;

-- name: UpdateEntries :exec
UPDATE entries
set account_id = $2,
    amount = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;
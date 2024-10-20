// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfer.sql

package db

import (
	"context"
	"database/sql"
)

const createTransfers = `-- name: CreateTransfers :one
INSERT INTO transfers (
    id, from_account_id, to_account_id, amount
) VALUES ($1, $2, $3, $4) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransfersParams struct {
	ID            int64         `json:"id"`
	FromAccountID sql.NullInt64 `json:"from_account_id"`
	ToAccountID   sql.NullInt64 `json:"to_account_id"`
	Amount        int64         `json:"amount"`
}

func (q *Queries) CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfers,
		arg.ID,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY from_account_id
`

func (q *Queries) ListTransfers(ctx context.Context) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfers = `-- name: UpdateTransfers :exec
UPDATE transfers
set from_account_id = $2,
    to_account_id = $3,
    amount = $4
WHERE id = $1
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type UpdateTransfersParams struct {
	ID            int64         `json:"id"`
	FromAccountID sql.NullInt64 `json:"from_account_id"`
	ToAccountID   sql.NullInt64 `json:"to_account_id"`
	Amount        int64         `json:"amount"`
}

func (q *Queries) UpdateTransfers(ctx context.Context, arg UpdateTransfersParams) error {
	_, err := q.db.ExecContext(ctx, updateTransfers,
		arg.ID,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	return err
}

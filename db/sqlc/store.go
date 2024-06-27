package db

import (
	"context"
	"database/sql"
	"fmt"
)

// provide all fxns to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// performs a money transfer from one account to the other
// it creates a transfer record, add account entries & update accounts balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	var err error
	var fromAccountID sql.NullInt64
	fromAccountID.Int64 = arg.FromAccountID
	fromAccountID.Valid = true

	var toAccountID sql.NullInt64
	toAccountID.Int64 = arg.ToAccountID
	toAccountID.Valid = true

	err = store.execTx(ctx, func(q *Queries) error {
		result.Transfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		return nil
	})
	return result, err
}

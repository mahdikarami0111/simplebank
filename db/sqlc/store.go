package db

import (
	"context"
	"database/sql"
)

type Store struct {
	Queries *Queries
	db      *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		tx.Rollback()
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transaction Transaction `json:"transaction"`
	FromAccount Account     `json:"from_account"`
	ToAccount   Account     `json:"to_account"`
	FromEntry   Entry       `json:"from_entry"`
	ToEntry     Entry       `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTX(ctx, func(q *Queries) error {
		var err error
		result.Transaction, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, err = q.AddAcountBalance(ctx, AddAcountBalanceParams{
				Amount: -1 * arg.Amount,
				ID:     arg.FromAccountID,
			})
			if err != nil {
				return err
			}

			result.ToAccount, err = q.AddAcountBalance(ctx, AddAcountBalanceParams{
				Amount: 1 * arg.Amount,
				ID:     arg.ToAccountID,
			})
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, err = q.AddAcountBalance(ctx, AddAcountBalanceParams{
				Amount: 1 * arg.Amount,
				ID:     arg.ToAccountID,
			})
			if err != nil {
				return err
			}
			result.FromAccount, err = q.AddAcountBalance(ctx, AddAcountBalanceParams{
				Amount: -1 * arg.Amount,
				ID:     arg.FromAccountID,
			})
			if err != nil {
				return err
			}

		}

		return nil

	})
	return result, err
}

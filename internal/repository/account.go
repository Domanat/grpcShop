package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	pool *pgxpool.Pool
}

func NewAccountRepository(
	pool *pgxpool.Pool,
) *AccountRepository {
	return &AccountRepository{pool: pool}
}

func (r *AccountRepository) Transfer(
	ctx context.Context,
	fromID int64,
	toID int64,
	amount int64,
) error {
	if fromID == toID {
		return errors.New("from == to error")
	}

	if amount < 1 {
		return errors.New("amount is less than 1")
	}

	query1 := `UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1`
	query2 := `UPDATE accounts SET balance = balance + $1 WHERE id = $2`

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction err: %w", err)
	}

	defer tx.Rollback(ctx)

	tag, err := tx.Exec(ctx, query1, amount, fromID)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	if tag.RowsAffected() != 1 {
		return errors.New("source account not found")
	}

	tag, err = tx.Exec(ctx, query2, amount, toID)
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}
	if tag.RowsAffected() != 1 {
		return errors.New("source account not found")
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

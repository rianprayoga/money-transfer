package repository

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrInsufucientBalance = errors.New("insuficient balance")
	ErrMerchantNotFound   = errors.New("source merchant not found")
)

type PgRepo struct {
	DB *sql.DB
}

func (pg *PgRepo) ReduceBalance(ctx context.Context, merchantId string, amount uint) error {
	tx, err := pg.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var isEnough bool
	if err = tx.
		QueryRowContext(ctx, `SELECT (balance >= $1) FROM merchants WHERE id = $2`, amount, merchantId).
		Scan(&isEnough); err != nil {
		if err == sql.ErrNoRows {
			return ErrMerchantNotFound
		}
		return err
	}

	if !isEnough {
		return ErrInsufucientBalance
	}

	var id uint
	if err = tx.QueryRowContext(
		ctx,
		`UPDATE merchants SET balance = balance - $1 WHERE id = $2 RETURNING id`, amount, merchantId).
		Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return ErrMerchantNotFound
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

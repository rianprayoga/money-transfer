package repository

import (
	"context"
	"database/sql"
	"errors"
	"moneytrx/internal/model"
)

var (
	ErrInsufucientBalance = errors.New("insuficient balance")
	ErrMerchantNotFound   = errors.New("source merchant not found")
)

type PgRepo struct {
	DB *sql.DB
}

func (pg *PgRepo) ReduceBalance(ctx context.Context, merchantId uint, amount uint) (*model.TrxRecord, error) {
	tx, err := pg.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var isEnough bool
	if err = tx.
		QueryRowContext(ctx, `SELECT (balance >= $1) FROM merchants WHERE id = $2 FOR UPDATE`, amount, merchantId).
		Scan(&isEnough); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMerchantNotFound
		}
		return nil, err
	}

	if !isEnough {
		return nil, ErrInsufucientBalance
	}

	var id uint
	if err = tx.QueryRowContext(
		ctx,
		`UPDATE merchants SET balance = balance - $1 WHERE id = $2 RETURNING id`, amount, merchantId).
		Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMerchantNotFound
		}

		return nil, err
	}

	var res model.TrxRecord
	if err = tx.QueryRowContext(
		ctx,
		`INSERT INTO transactions(merchant_id, amount) VALUES($1,$2) RETURNING id, amount`, merchantId, amount).
		Scan(&res.Id, &res.Amount); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMerchantNotFound
		}
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &res, nil
}

func (pg *PgRepo) SuccessTrx(ctx context.Context, trxId uint) error {
	tx, err := pg.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var id uint
	if err = tx.QueryRowContext(
		ctx,
		`UPDATE transactions SET status= 'SUCCESS' WHERE id = $1 RETURNING id`, trxId).
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

func (pg *PgRepo) FailedTrx(ctx context.Context, trxId uint, merchantId uint, amount uint) error {
	tx, err := pg.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var id uint
	if err = tx.QueryRowContext(
		ctx,
		`UPDATE merchants SET balance = balance + $1 WHERE id = $2 RETURNING id`, amount, merchantId).
		Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return ErrMerchantNotFound
		}
		return err
	}

	if err = tx.QueryRowContext(
		ctx,
		`UPDATE transactions SET status= 'FAILED' WHERE id = $1 RETURNING id`, trxId).
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

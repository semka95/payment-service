// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: queries.sql

package repository

import (
	"context"

	"github.com/shopspring/decimal"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments(
    user_id, email, amount, currency, payment_status
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, user_id, email, amount, currency, payment_status, created_at, updated_at
`

type CreatePaymentParams struct {
	UserID        int64           `json:"user_id"`
	Email         string          `json:"email"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      ValidCurrency   `json:"currency"`
	PaymentStatus ValidStatus     `json:"payment_status"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.UserID,
		arg.Email,
		arg.Amount,
		arg.Currency,
		arg.PaymentStatus,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Email,
		&i.Amount,
		&i.Currency,
		&i.PaymentStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const discardPayment = `-- name: DiscardPayment :execrows
DELETE FROM payments 
WHERE id = $1 AND payment_status NOT IN ('success', 'failure')
`

func (q *Queries) DiscardPayment(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, discardPayment, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getPaymentStatusByID = `-- name: GetPaymentStatusByID :one
SELECT payment_status FROM payments
WHERE id = $1
`

func (q *Queries) GetPaymentStatusByID(ctx context.Context, id int64) (ValidStatus, error) {
	row := q.db.QueryRowContext(ctx, getPaymentStatusByID, id)
	var payment_status ValidStatus
	err := row.Scan(&payment_status)
	return payment_status, err
}

const listUserPaymentsByEmail = `-- name: ListUserPaymentsByEmail :many
SELECT id, user_id, email, amount, currency, payment_status, created_at, updated_at FROM payments
WHERE email = $1 AND id > $2
LIMIT $3
`

type ListUserPaymentsByEmailParams struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
	Limit int32  `json:"limit"`
}

func (q *Queries) ListUserPaymentsByEmail(ctx context.Context, arg ListUserPaymentsByEmailParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listUserPaymentsByEmail, arg.Email, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Email,
			&i.Amount,
			&i.Currency,
			&i.PaymentStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listUserPaymentsByID = `-- name: ListUserPaymentsByID :many
SELECT id, user_id, email, amount, currency, payment_status, created_at, updated_at FROM payments
WHERE user_id = $1 AND id > $2
LIMIT $3
`

type ListUserPaymentsByIDParams struct {
	UserID int64 `json:"user_id"`
	ID     int64 `json:"id"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListUserPaymentsByID(ctx context.Context, arg ListUserPaymentsByIDParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listUserPaymentsByID, arg.UserID, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Email,
			&i.Amount,
			&i.Currency,
			&i.PaymentStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updatePaymentStatus = `-- name: UpdatePaymentStatus :execrows
UPDATE payments
SET payment_status = $2,
    updated_at = NOW()
WHERE id = $1 AND payment_status NOT IN ('success', 'failure')
`

type UpdatePaymentStatusParams struct {
	ID            int64       `json:"id"`
	PaymentStatus ValidStatus `json:"payment_status"`
}

func (q *Queries) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updatePaymentStatus, arg.ID, arg.PaymentStatus)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

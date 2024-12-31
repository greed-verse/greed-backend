// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: wallets.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createWallet = `-- name: CreateWallet :one
INSERT INTO wallets (user_id, balance)
VALUES ($1, $2)
RETURNING id, user_id, balance, updated_at
`

type CreateWalletParams struct {
	UserID  string         `json:"user_id"`
	Balance pgtype.Numeric `json:"balance"`
}

func (q *Queries) CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error) {
	row := q.db.QueryRow(ctx, createWallet, arg.UserID, arg.Balance)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteWallet = `-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1
`

func (q *Queries) DeleteWallet(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteWallet, id)
	return err
}

const getWalletByID = `-- name: GetWalletByID :one
SELECT id, user_id, balance, updated_at
FROM wallets
WHERE id = $1
`

func (q *Queries) GetWalletByID(ctx context.Context, id string) (Wallet, error) {
	row := q.db.QueryRow(ctx, getWalletByID, id)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.UpdatedAt,
	)
	return i, err
}

const getWalletByUserID = `-- name: GetWalletByUserID :one
SELECT id, user_id, balance, updated_at
FROM wallets
WHERE user_id = $1
`

func (q *Queries) GetWalletByUserID(ctx context.Context, userID string) (Wallet, error) {
	row := q.db.QueryRow(ctx, getWalletByUserID, userID)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.UpdatedAt,
	)
	return i, err
}

const listWallets = `-- name: ListWallets :many
SELECT id, user_id, balance, updated_at
FROM wallets
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2
`

type ListWalletsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListWallets(ctx context.Context, arg ListWalletsParams) ([]Wallet, error) {
	rows, err := q.db.Query(ctx, listWallets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Wallet
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Balance,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateWalletBalance = `-- name: UpdateWalletBalance :one
UPDATE wallets
SET balance = $1, updated_at = NOW()
WHERE user_id = $2
RETURNING id, user_id, balance, updated_at
`

type UpdateWalletBalanceParams struct {
	Balance pgtype.Numeric `json:"balance"`
	UserID  string         `json:"user_id"`
}

func (q *Queries) UpdateWalletBalance(ctx context.Context, arg UpdateWalletBalanceParams) (Wallet, error) {
	row := q.db.QueryRow(ctx, updateWalletBalance, arg.Balance, arg.UserID)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.UpdatedAt,
	)
	return i, err
}

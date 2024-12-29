-- name: CreateWallet :one
INSERT INTO wallets (user_id, balance)
VALUES ($1, $2)
RETURNING id, user_id, balance, updated_at;

-- name: GetWalletByID :one
SELECT id, user_id, balance, updated_at
FROM wallets
WHERE id = $1;

-- name: GetWalletByUserID :one
SELECT id, user_id, balance, updated_at
FROM wallets
WHERE user_id = $1;

-- name: UpdateWalletBalance :one
UPDATE wallets
SET balance = $1, updated_at = NOW()
WHERE user_id = $2
RETURNING id, user_id, balance, updated_at;

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1;

-- name: ListWallets :many
SELECT id, user_id, balance, updated_at
FROM wallets
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2;

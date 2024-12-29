-- name: CreateUser :one
INSERT INTO users (username, email)
VALUES ($1, $2)
RETURNING id, username, email, created_at;

-- name: GetUserByID :one
SELECT id, username, email, created_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, username, email, created_at
FROM users
WHERE email = $1;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $1
WHERE id = $2
RETURNING id, username, email, created_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, username, email, created_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CheckEmailExists :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
) AS exists;

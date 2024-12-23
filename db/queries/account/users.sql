-- name: CreateUser :one
INSERT INTO users (email, name, first_name, profile_image )
VALUES ($1, $2, $3, $4)
RETURNING id, email, name, first_name, profile_image, created_at;

-- name: GetUserByID :one
SELECT id, email, name, first_name, profile_image, created_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, name, first_name, profile_image, created_at
FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT id, email, name, first_name, profile_image, created_at
FROM users
ORDER BY created_at DESC;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;

-- name: CheckEmailExists :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
) AS exists;

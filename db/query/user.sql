-- name: CreateUser :one
INSERT INTO users (
    user_name,
    PASSWORD,
    full_name)
VALUES (
    $1,
    $2,
    $3)
RETURNING
    *;

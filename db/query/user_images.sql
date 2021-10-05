-- query.sql

-- name: CreateUserImage :one
INSERT INTO user_images (
    UserId,
    Url
)
VALUES (
   $1,
   $2
    )
    RETURNING
    *;
-- name: ListImage :many
SELECT * FROM user_images
WHERE UserId = $1;
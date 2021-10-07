-- query.sql

-- name: CreateUserImage :one
INSERT INTO user_images (
    user_id,
    url,
    type
)
VALUES (
   $1,
   $2,
   $3
    )
    RETURNING
    *;
-- name: ListImage :many
SELECT * FROM user_images
WHERE user_id = $1;
-- query.sql

-- name: CreateUserRole :one
INSERT INTO user_role(
    user_id,
    role_id
)
VALUES (
    $1,
    $2
       )
    RETURNING
    *;
-- name: GetRoleByUserId :one
SELECT role_id FROM user_role
WHERE user_id = $1 LIMIT 1;
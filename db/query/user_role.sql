-- query.sql

-- name: CreateUserRole :one
INSERT INTO user_role(
    UserId,
    RoleId
)
VALUES (
    $1,
    $2
       )
    RETURNING
    *;
-- name: GetRoleByUserId :one
SELECT RoleId FROM user_role
WHERE UserId = $1 LIMIT 1;
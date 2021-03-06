// Code generated by sqlc. DO NOT EDIT.
// source: user_role.sql

package db

import (
	"context"
)

const createUserRole = `-- name: CreateUserRole :one

INSERT INTO user_role(
    user_id,
    role_id
)
VALUES (
    $1,
    $2
       )
    RETURNING
    id, user_id, role_id
`

type CreateUserRoleParams struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

// query.sql
func (q *Queries) CreateUserRole(ctx context.Context, arg CreateUserRoleParams) (UserRole, error) {
	row := q.db.QueryRowContext(ctx, createUserRole, arg.UserID, arg.RoleID)
	var i UserRole
	err := row.Scan(&i.ID, &i.UserID, &i.RoleID)
	return i, err
}

const getRoleByUserId = `-- name: GetRoleByUserId :one
SELECT role_id FROM user_role
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetRoleByUserId(ctx context.Context, userID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getRoleByUserId, userID)
	var role_id int64
	err := row.Scan(&role_id)
	return role_id, err
}

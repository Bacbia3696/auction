// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    user_name,
    password,
    full_name,
    email,
    address,
    phone,
    birthdate,
    id_card,
    id_card_address,
    id_card_date,
    bank_id,
    bank_owner,
    bank_name,
    status,
    organization_name,
    organization_id ,
    organization_date ,
    organization_address,
    created_at,
    updated_at
)
VALUES (
       $1,
       $2,
       $3,
       $4,
       $5,
       $6,
       $7,
       $8,
       $9,
       $10,
       $11,
       $12,
       $13,
       $14,
       $15,
       $16,
       $17,
       $18,
       $19,
       $20)
    RETURNING
    id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at
`

type CreateUserParams struct {
	UserName            string         `json:"user_name"`
	Password            string         `json:"password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	Address             string         `json:"address"`
	Phone               string         `json:"phone"`
	Birthdate           sql.NullTime   `json:"birthdate"`
	IDCard              string         `json:"id_card"`
	IDCardAddress       string         `json:"id_card_address"`
	IDCardDate          time.Time      `json:"id_card_date"`
	BankID              string         `json:"bank_id"`
	BankOwner           string         `json:"bank_owner"`
	BankName            string         `json:"bank_name"`
	Status              int32          `json:"status"`
	OrganizationName    sql.NullString `json:"organization_name"`
	OrganizationID      sql.NullString `json:"organization_id"`
	OrganizationDate    sql.NullTime   `json:"organization_date"`
	OrganizationAddress sql.NullString `json:"organization_address"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
}

// query.sql
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserName,
		arg.Password,
		arg.FullName,
		arg.Email,
		arg.Address,
		arg.Phone,
		arg.Birthdate,
		arg.IDCard,
		arg.IDCardAddress,
		arg.IDCardDate,
		arg.BankID,
		arg.BankOwner,
		arg.BankName,
		arg.Status,
		arg.OrganizationName,
		arg.OrganizationID,
		arg.OrganizationDate,
		arg.OrganizationAddress,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getByEmail = `-- name: GetByEmail :one
SELECT id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getById = `-- name: GetById :one
SELECT id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getByIdCard = `-- name: GetByIdCard :one
SELECT id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at FROM users
WHERE id_card = $1 LIMIT 1
`

func (q *Queries) GetByIdCard(ctx context.Context, idCard string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByIdCard, idCard)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getByUserName = `-- name: GetByUserName :one
SELECT id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at FROM users
WHERE user_name = $1 LIMIT 1
`

func (q *Queries) GetByUserName(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByUserName, userName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getByUserNameActive = `-- name: GetByUserNameActive :one
SELECT id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at FROM users
WHERE user_name = $1 AND status > 0 LIMIT 1
`

func (q *Queries) GetByUserNameActive(ctx context.Context, userName string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByUserNameActive, userName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListUser = `-- name: GetListUser :many
SELECT id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at FROM users
WHERE ( user_name LIKE  $1 OR full_name LIKE  $1 OR organization_name LIKE  $1 OR id_card LIKE  $1 OR organization_id LIKE  $1 OR email  LIKE  $1)
ORDER BY id ASC LIMIT $3 OFFSET $2
`

type GetListUserParams struct {
	UserName string `json:"user_name"`
	Offset   int32  `json:"offset"`
	Limit    int32  `json:"limit"`
}

func (q *Queries) GetListUser(ctx context.Context, arg GetListUserParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getListUser, arg.UserName, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.UserName,
			&i.Password,
			&i.FullName,
			&i.Email,
			&i.Address,
			&i.Phone,
			&i.Birthdate,
			&i.IDCard,
			&i.IDCardAddress,
			&i.IDCardDate,
			&i.BankID,
			&i.BankOwner,
			&i.BankName,
			&i.Status,
			&i.OrganizationName,
			&i.OrganizationID,
			&i.OrganizationDate,
			&i.OrganizationAddress,
			&i.Position,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalUser = `-- name: GetTotalUser :one
SELECT COUNT(*) FROM users
WHERE ( user_name LIKE  $1 OR full_name LIKE  $1 OR organization_name LIKE  $1 OR id_card LIKE  $1 OR organization_id LIKE  $1 OR email  LIKE  $1)
`

func (q *Queries) GetTotalUser(ctx context.Context, userName string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalUser, userName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateStatus = `-- name: UpdateStatus :one
UPDATE users
SET status = $1
WHERE  id = $2
    RETURNING
    id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at
`

type UpdateStatusParams struct {
	Status int32 `json:"status"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateStatus, arg.Status, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Password,
		&i.FullName,
		&i.Email,
		&i.Address,
		&i.Phone,
		&i.Birthdate,
		&i.IDCard,
		&i.IDCardAddress,
		&i.IDCardDate,
		&i.BankID,
		&i.BankOwner,
		&i.BankName,
		&i.Status,
		&i.OrganizationName,
		&i.OrganizationID,
		&i.OrganizationDate,
		&i.OrganizationAddress,
		&i.Position,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

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
    organization_address
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
       $18
       )
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

const getAllListUserBidAuction = `-- name: GetAllListUserBidAuction :many
SELECT u.user_name, u.full_name, u.phone, u.email, u.id_card, u.bank_id, b.price, b.created_at
FROM bid as b
         INNER JOIN users as u ON b.user_id = u.id
WHERE b.auction_id = $1
ORDER BY b.price DESC  LIMIT $3 OFFSET $2
`

type GetAllListUserBidAuctionParams struct {
	AuctionID int32 `json:"auction_id"`
	Offset    int32 `json:"offset"`
	Limit     int32 `json:"limit"`
}

type GetAllListUserBidAuctionRow struct {
	UserName  string    `json:"user_name"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IDCard    string    `json:"id_card"`
	BankID    string    `json:"bank_id"`
	Price     int32     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) GetAllListUserBidAuction(ctx context.Context, arg GetAllListUserBidAuctionParams) ([]GetAllListUserBidAuctionRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllListUserBidAuction, arg.AuctionID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllListUserBidAuctionRow{}
	for rows.Next() {
		var i GetAllListUserBidAuctionRow
		if err := rows.Scan(
			&i.UserName,
			&i.FullName,
			&i.Phone,
			&i.Email,
			&i.IDCard,
			&i.BankID,
			&i.Price,
			&i.CreatedAt,
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

const getAllListUserRegisterAuction = `-- name: GetAllListUserRegisterAuction :many
SELECT u.user_name, u.full_name, u.phone, u.email, u.id_card, u.bank_id, ra.created_at, ra.status as verify
FROM register_auction as ra
INNER JOIN auctions as au ON ra.auction_id = au.id
INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1
ORDER BY u.id ASC LIMIT $3 OFFSET $2
`

type GetAllListUserRegisterAuctionParams struct {
	AuctionID int32 `json:"auction_id"`
	Offset    int32 `json:"offset"`
	Limit     int32 `json:"limit"`
}

type GetAllListUserRegisterAuctionRow struct {
	UserName  string    `json:"user_name"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IDCard    string    `json:"id_card"`
	BankID    string    `json:"bank_id"`
	CreatedAt time.Time `json:"created_at"`
	Verify    int32     `json:"verify"`
}

func (q *Queries) GetAllListUserRegisterAuction(ctx context.Context, arg GetAllListUserRegisterAuctionParams) ([]GetAllListUserRegisterAuctionRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllListUserRegisterAuction, arg.AuctionID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllListUserRegisterAuctionRow{}
	for rows.Next() {
		var i GetAllListUserRegisterAuctionRow
		if err := rows.Scan(
			&i.UserName,
			&i.FullName,
			&i.Phone,
			&i.Email,
			&i.IDCard,
			&i.BankID,
			&i.CreatedAt,
			&i.Verify,
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

const getListUserRegisterAuctionByStatus = `-- name: GetListUserRegisterAuctionByStatus :many
SELECT u.user_name, u.full_name, u.phone, u.email, u.id_card, u.bank_id, ra.created_at, ra.status as verify
FROM register_auction as ra
         INNER JOIN auctions as au ON ra.auction_id = au.id
         INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1 AND ra.status =$2
ORDER BY u.id ASC LIMIT $4 OFFSET $3
`

type GetListUserRegisterAuctionByStatusParams struct {
	AuctionID int32 `json:"auction_id"`
	Status    int32 `json:"status"`
	Offset    int32 `json:"offset"`
	Limit     int32 `json:"limit"`
}

type GetListUserRegisterAuctionByStatusRow struct {
	UserName  string    `json:"user_name"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IDCard    string    `json:"id_card"`
	BankID    string    `json:"bank_id"`
	CreatedAt time.Time `json:"created_at"`
	Verify    int32     `json:"verify"`
}

func (q *Queries) GetListUserRegisterAuctionByStatus(ctx context.Context, arg GetListUserRegisterAuctionByStatusParams) ([]GetListUserRegisterAuctionByStatusRow, error) {
	rows, err := q.db.QueryContext(ctx, getListUserRegisterAuctionByStatus,
		arg.AuctionID,
		arg.Status,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListUserRegisterAuctionByStatusRow{}
	for rows.Next() {
		var i GetListUserRegisterAuctionByStatusRow
		if err := rows.Scan(
			&i.UserName,
			&i.FullName,
			&i.Phone,
			&i.Email,
			&i.IDCard,
			&i.BankID,
			&i.CreatedAt,
			&i.Verify,
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

const getTotalListUserBidAuction = `-- name: GetTotalListUserBidAuction :one
SELECT  COUNT(*)
FROM bid as b
         INNER JOIN users as u ON b.user_id = u.id
WHERE b.auction_id = $1
`

func (q *Queries) GetTotalListUserBidAuction(ctx context.Context, auctionID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalListUserBidAuction, auctionID)
	var count int64
	err := row.Scan(&count)
	return count, err
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

const getTotalUserRegisterAuction = `-- name: GetTotalUserRegisterAuction :one
SELECT  COUNT(*)
FROM register_auction as ra
         INNER JOIN auctions as au ON ra.auction_id = au.id
         INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1
`

func (q *Queries) GetTotalUserRegisterAuction(ctx context.Context, auctionID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalUserRegisterAuction, auctionID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTotalUserRegisterAuctionByStatus = `-- name: GetTotalUserRegisterAuctionByStatus :one
SELECT  COUNT(*)
FROM register_auction as ra
         INNER JOIN auctions as au ON ra.auction_id = au.id
         INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1 AND ra.status =$2
`

type GetTotalUserRegisterAuctionByStatusParams struct {
	AuctionID int32 `json:"auction_id"`
	Status    int32 `json:"status"`
}

func (q *Queries) GetTotalUserRegisterAuctionByStatus(ctx context.Context, arg GetTotalUserRegisterAuctionByStatusParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalUserRegisterAuctionByStatus, arg.AuctionID, arg.Status)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE users
SET password = $1
WHERE  id = $2
    RETURNING
    id, user_name, password, full_name, email, address, phone, birthdate, id_card, id_card_address, id_card_date, bank_id, bank_owner, bank_name, status, organization_name, organization_id, organization_date, organization_address, position, created_at, updated_at
`

type UpdatePasswordParams struct {
	Password string `json:"password"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.Password, arg.ID)
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

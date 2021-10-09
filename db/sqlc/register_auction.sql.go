// Code generated by sqlc. DO NOT EDIT.
// source: register_auction.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createRegisterAuction = `-- name: CreateRegisterAuction :one
INSERT INTO register_auction (
    auction_id,
    user_id,
    status,
    updated_at,
    created_at
)
VALUES (
       $1,
       $2,
       $3,
       $4,
       $5
       )
    RETURNING
    id, auction_id, user_id, status, updated_at, created_at
`

type CreateRegisterAuctionParams struct {
	AuctionID int32        `json:"auction_id"`
	UserID    int32        `json:"user_id"`
	Status    int32        `json:"status"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
}

// query.sql
func (q *Queries) CreateRegisterAuction(ctx context.Context, arg CreateRegisterAuctionParams) (RegisterAuction, error) {
	row := q.db.QueryRowContext(ctx, createRegisterAuction,
		arg.AuctionID,
		arg.UserID,
		arg.Status,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	var i RegisterAuction
	err := row.Scan(
		&i.ID,
		&i.AuctionID,
		&i.UserID,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getListRegisterAuction = `-- name: GetListRegisterAuction :many
SELECT u.id, u.user_name, u.password, u.full_name, u.email, u.address, u.phone, u.birthdate, u.id_card, u.id_card_address, u.id_card_date, u.bank_id, u.bank_owner, u.bank_name, u.status, u.organization_name, u.organization_id, u.organization_date, u.organization_address, u.position, u.created_at, u.updated_at FROM register_auction as ra INNER JOIN users as u ON ra.user_id = u.id
WHERE  ra.auction_id = $1
ORDER BY id ASC LIMIT $3 OFFSET $2
`

type GetListRegisterAuctionParams struct {
	AuctionID int32 `json:"auction_id"`
	Offset    int32 `json:"offset"`
	Limit     int32 `json:"limit"`
}

func (q *Queries) GetListRegisterAuction(ctx context.Context, arg GetListRegisterAuctionParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getListRegisterAuction, arg.AuctionID, arg.Offset, arg.Limit)
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

const getListRegisterAuctionByUserId = `-- name: GetListRegisterAuctionByUserId :many
SELECT au.id, au.code, au.owner, au.organization, au.register_start_date, au.register_end_date, au.bid_start_date, au.bid_end_date, au.start_price, au.status, au.updated_at, au.created_at, ra.status as verify FROM register_auction as ra INNER JOIN auctions as au ON ra.auction_id = au.id
WHERE ra.user_id = $1 ORDER BY id ASC LIMIT $3 OFFSET $2
`

type GetListRegisterAuctionByUserIdParams struct {
	UserID int32 `json:"user_id"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type GetListRegisterAuctionByUserIdRow struct {
	ID                int32        `json:"id"`
	Code              string       `json:"code"`
	Owner             string       `json:"owner"`
	Organization      string       `json:"organization"`
	RegisterStartDate time.Time    `json:"register_start_date"`
	RegisterEndDate   time.Time    `json:"register_end_date"`
	BidStartDate      time.Time    `json:"bid_start_date"`
	BidEndDate        time.Time    `json:"bid_end_date"`
	StartPrice        int32        `json:"start_price"`
	Status            int32        `json:"status"`
	UpdatedAt         sql.NullTime `json:"updated_at"`
	CreatedAt         time.Time    `json:"created_at"`
	Verify            int32        `json:"verify"`
}

func (q *Queries) GetListRegisterAuctionByUserId(ctx context.Context, arg GetListRegisterAuctionByUserIdParams) ([]GetListRegisterAuctionByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getListRegisterAuctionByUserId, arg.UserID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListRegisterAuctionByUserIdRow{}
	for rows.Next() {
		var i GetListRegisterAuctionByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Owner,
			&i.Organization,
			&i.RegisterStartDate,
			&i.RegisterEndDate,
			&i.BidStartDate,
			&i.BidEndDate,
			&i.StartPrice,
			&i.Status,
			&i.UpdatedAt,
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

const getRegisterAuctionById = `-- name: GetRegisterAuctionById :one
SELECT id, auction_id, user_id, status, updated_at, created_at FROM register_auction
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRegisterAuctionById(ctx context.Context, id int32) (RegisterAuction, error) {
	row := q.db.QueryRowContext(ctx, getRegisterAuctionById, id)
	var i RegisterAuction
	err := row.Scan(
		&i.ID,
		&i.AuctionID,
		&i.UserID,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getRegisterAuctionByUserId = `-- name: GetRegisterAuctionByUserId :one
SELECT au.id, au.code, au.owner, au.organization, au.register_start_date, au.register_end_date, au.bid_start_date, au.bid_end_date, au.start_price, au.status, au.updated_at, au.created_at FROM register_auction as ra INNER JOIN auctions as au ON ra.auction_id = au.id
WHERE ra.user_id = $1 AND ra.auction_id = $2
`

type GetRegisterAuctionByUserIdParams struct {
	UserID    int32 `json:"user_id"`
	AuctionID int32 `json:"auction_id"`
}

func (q *Queries) GetRegisterAuctionByUserId(ctx context.Context, arg GetRegisterAuctionByUserIdParams) (Auction, error) {
	row := q.db.QueryRowContext(ctx, getRegisterAuctionByUserId, arg.UserID, arg.AuctionID)
	var i Auction
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Owner,
		&i.Organization,
		&i.RegisterStartDate,
		&i.RegisterEndDate,
		&i.BidStartDate,
		&i.BidEndDate,
		&i.StartPrice,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getTotalRegisterAuction = `-- name: GetTotalRegisterAuction :one
SELECT COUNT(*) FROM register_auction
WHERE  auction_id = $1
`

func (q *Queries) GetTotalRegisterAuction(ctx context.Context, auctionID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalRegisterAuction, auctionID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateStatusRegisterAuction = `-- name: UpdateStatusRegisterAuction :one
UPDATE register_auction
SET status = $1
WHERE  id = $2
    RETURNING
    id, auction_id, user_id, status, updated_at, created_at
`

type UpdateStatusRegisterAuctionParams struct {
	Status int32 `json:"status"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateStatusRegisterAuction(ctx context.Context, arg UpdateStatusRegisterAuctionParams) (RegisterAuction, error) {
	row := q.db.QueryRowContext(ctx, updateStatusRegisterAuction, arg.Status, arg.ID)
	var i RegisterAuction
	err := row.Scan(
		&i.ID,
		&i.AuctionID,
		&i.UserID,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

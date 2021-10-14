// Code generated by sqlc. DO NOT EDIT.
// source: bid.sql

package db

import (
	"context"
	"time"
)

const createBid = `-- name: CreateBid :one
INSERT INTO bid (
    auction_id,
    user_id,
    price,
    status,
    updated_at,
    created_at
)
VALUES (
       $1,
       $2,
       $3,
       $4,
       $5,
       $6
    )
    RETURNING
    id, auction_id, user_id, price, status, updated_at, created_at
`

type CreateBidParams struct {
	AuctionID int32     `json:"auction_id"`
	UserID    int32     `json:"user_id"`
	Price     int32     `json:"price"`
	Status    int32     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// query.sql
func (q *Queries) CreateBid(ctx context.Context, arg CreateBidParams) (Bid, error) {
	row := q.db.QueryRowContext(ctx, createBid,
		arg.AuctionID,
		arg.UserID,
		arg.Price,
		arg.Status,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	var i Bid
	err := row.Scan(
		&i.ID,
		&i.AuctionID,
		&i.UserID,
		&i.Price,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getMaxBid = `-- name: GetMaxBid :one
SELECT MAX(price) FROM bid
WHERE auction_id = $1
`

func (q *Queries) GetMaxBid(ctx context.Context, auctionID int32) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, getMaxBid, auctionID)
	var max interface{}
	err := row.Scan(&max)
	return max, err
}

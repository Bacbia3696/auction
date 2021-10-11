-- query.sql
-- name: CreateBid :one
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
    *;
-- name: GetMaxBid :one
SELECT MAX(price) FROM bid
WHERE auction_id = $1;
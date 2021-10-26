-- query.sql
-- name: CreateBid :one
INSERT INTO bid (
    auction_id,
    user_id,
    price,
    status
)
VALUES (
       $1,
       $2,
       $3,
       $4
    )
    RETURNING
    *;
-- name: GetMaxBid :one
SELECT MAX(price) FROM bid
WHERE auction_id = $1;


-- name: GetTotalUserBid :one
SELECT COUNT(*) FROM bid
WHERE auction_id = $1 AND user_id=$2;
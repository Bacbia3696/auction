-- query.sql
-- name: CreateRegisterAuction :one
INSERT INTO register_auction (
    auction_id,
    user_id,
    status
)
VALUES (
       $1,
       $2,
       $3
       )
    RETURNING
    *;
-- name: UpdateStatusRegisterAuction :one
UPDATE register_auction
SET status = $1
WHERE  id = $2
    RETURNING
    *;

-- name: GetRegisterAuctionById :one
SELECT * FROM register_auction
WHERE id = $1 LIMIT 1;

-- name: GetListRegisterAuctionByUserId :many
SELECT au.*, ra.status as verify FROM register_auction as ra INNER JOIN auctions as au ON ra.auction_id = au.id
WHERE ra.user_id = $1 ORDER BY id ASC LIMIT $3 OFFSET $2;;

-- name: GetRegisterAuctionByUserId :one
SELECT au.*, ra.status as verify FROM register_auction as ra INNER JOIN auctions as au ON ra.auction_id = au.id
WHERE ra.user_id = $1 AND ra.auction_id = $2;

-- name: GetListRegisterAuction :many
SELECT u.* FROM register_auction as ra INNER JOIN users as u ON ra.user_id = u.id
WHERE  ra.auction_id = $1
ORDER BY id ASC LIMIT $3 OFFSET $2;

-- name: GetTotalRegisterAuction :one
SELECT COUNT(*) FROM register_auction
WHERE  auction_id = $1
;

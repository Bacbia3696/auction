-- query.sql
-- name: CreateAuction :one
INSERT INTO auctions (
    code,
    owner,
    organization,
    register_start_date,
    register_end_date,
    bid_start_date,
    bid_end_date,
    start_price,
    status,
    type ,
    updated_at,
    created_at
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
       $12
         )
    RETURNING
    *;
-- name: UpdateStatusAuction :one
UPDATE auctions
SET status = $1
WHERE  id = $2
    RETURNING
    *;
-- name: GetByCode :one
SELECT * FROM auctions
WHERE code = $1 LIMIT 1;

-- name: GetAuctionById :one
SELECT * FROM auctions
WHERE id = $1 LIMIT 1;

-- name: GetListAuction :many
SELECT * FROM auctions
WHERE ( code LIKE  $1 OR owner LIKE  $1 OR organization LIKE  $1 )
ORDER BY id ASC LIMIT $3 OFFSET $2;

-- name: GetTotalAuction :one
SELECT COUNT(*) FROM auctions
WHERE ( code LIKE  $1 OR owner LIKE  $1 OR organization LIKE  $1)
;
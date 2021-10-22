-- query.sql
-- name: CreateUser :one
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
    *;

-- name: GetById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetByUserName :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- name: GetByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetByIdCard :one
SELECT * FROM users
WHERE id_card = $1 LIMIT 1;

-- name: GetByUserNameActive :one
SELECT * FROM users
WHERE user_name = $1 AND status > 0 LIMIT 1;

-- name: UpdateStatus :one
UPDATE users
SET status = $1
WHERE  id = $2
    RETURNING
    *;
-- name: GetListUser :many
SELECT * FROM users
WHERE ( user_name LIKE  $1 OR full_name LIKE  $1 OR organization_name LIKE  $1 OR id_card LIKE  $1 OR organization_id LIKE  $1 OR email  LIKE  $1)
ORDER BY id ASC LIMIT $3 OFFSET $2;

-- name: GetTotalUser :one
SELECT COUNT(*) FROM users
WHERE ( user_name LIKE  $1 OR full_name LIKE  $1 OR organization_name LIKE  $1 OR id_card LIKE  $1 OR organization_id LIKE  $1 OR email  LIKE  $1)
;
-- name: UpdatePassword :one
UPDATE users
SET password = $1
WHERE  id = $2
    RETURNING
    *;


-- name: GetAllListUserRegisterAuction :many
SELECT u.user_name, u.full_name, u.phone, u.email, u.id_card, u.bank_id, ra.created_at, ra.status as verify
FROM register_auction as ra
INNER JOIN auctions as au ON ra.auction_id = au.id
INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1
ORDER BY u.id ASC LIMIT $3 OFFSET $2;

-- name: GetListUserRegisterAuctionByStatus :many
SELECT u.user_name, u.full_name, u.phone, u.email, u.id_card, u.bank_id, ra.created_at, ra.status as verify
FROM register_auction as ra
         INNER JOIN auctions as au ON ra.auction_id = au.id
         INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1 AND ra.status =$2
ORDER BY u.id ASC LIMIT $4 OFFSET $3;

-- name: GetTotalUserRegisterAuction :one
SELECT  COUNT(*)
FROM register_auction as ra
         INNER JOIN auctions as au ON ra.auction_id = au.id
         INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1;

-- name: GetTotalUserRegisterAuctionByStatus :one
SELECT  COUNT(*)
FROM register_auction as ra
         INNER JOIN auctions as au ON ra.auction_id = au.id
         INNER JOIN users as u ON ra.user_id = u.id
WHERE ra.auction_id = $1 AND ra.status =$2;

-- name: GetAllListUserBidAuction :many
SELECT u.user_name, u.full_name, u.phone, u.email, u.id_card, u.bank_id, b.price, b.created_at
FROM bid as b
         INNER JOIN users as u ON b.user_id = u.id
WHERE b.auction_id = $1
ORDER BY b.id DESC  LIMIT $3 OFFSET $2;

-- name: GetTotalListUserBidAuction :one
SELECT  COUNT(*)
FROM bid as b
         INNER JOIN users as u ON b.user_id = u.id
WHERE b.auction_id = $1;
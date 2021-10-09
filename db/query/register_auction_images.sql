-- query.sql

-- name: CreateRegisterAuctionImage :one
INSERT INTO register_auction_images (
    register_auction_id,
    url
)
VALUES (
           $1,
           $2
       )
    RETURNING
    *;
-- name: ListRegisterAuctionImage :many
SELECT * FROM register_auction_images
WHERE register_auction_id = $1;
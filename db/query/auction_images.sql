-- query.sql
-- name: CreateAuctionImage :one
INSERT INTO auction_images (
    auction_id,
    url)
VALUES (
    $1,
    $2)
RETURNING
    *;

-- name: ListAuctionImage :many
SELECT
    auction_images.url
FROM
    auction_images
WHERE
    auction_id = $1;

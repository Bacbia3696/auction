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
    created_at,
    updated_at
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
    $16)
RETURNING
    *;

-- name: CreateUserOrganization :one
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
    organization_address,
    created_at,
    updated_at
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
       $18,
       $19,
       $20))
    RETURNING
    *;

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



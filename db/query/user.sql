-- query.sql

-- name: CreateUser :one
INSERT INTO users (
    UserName,
    Password,
    FullName,
    Email,
    Address,
    Phone,
    BirthDate,
    IdCard,
    IdCardAddress,
    IdCardDate,
    BankId,
    BankOwner,
    BankName,
    Status,
    CreatedAt,
    UpdatedAt
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

-- name: GetByUserName :one
SELECT * FROM users
WHERE UserName = $1 LIMIT 1;

-- name: GetByEmail :one
SELECT * FROM users
WHERE Email = $1 LIMIT 1;

-- name: GetByIdCard :one
SELECT * FROM users
WHERE IdCard = $1 LIMIT 1;

-- name: GetByUserNameActive :one
SELECT * FROM users
WHERE UserName = $1 AND Status = 1 LIMIT 1;



// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Auction struct {
	ID                int32        `json:"id"`
	Code              string       `json:"code"`
	Owner             string       `json:"owner"`
	Organization      string       `json:"organization"`
	RegisterStartDate time.Time    `json:"register_start_date"`
	RegisterEndDate   time.Time    `json:"register_end_date"`
	BidStartDate      time.Time    `json:"bid_start_date"`
	BidEndDate        time.Time    `json:"bid_end_date"`
	StartPrice        int32        `json:"start_price"`
	Status            int32        `json:"status"`
	Type              int32        `json:"type"`
	UpdatedAt         sql.NullTime `json:"updated_at"`
	CreatedAt         time.Time    `json:"created_at"`
}

type AuctionImage struct {
	ID        int32  `json:"id"`
	AuctionID int32  `json:"auction_id"`
	Url       string `json:"url"`
}

type Bid struct {
	ID        int32     `json:"id"`
	AuctionID int32     `json:"auction_id"`
	UserID    int32     `json:"user_id"`
	Price     int32     `json:"price"`
	Status    int32     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterAuction struct {
	ID        int32        `json:"id"`
	AuctionID int32        `json:"auction_id"`
	UserID    int32        `json:"user_id"`
	Status    int32        `json:"status"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
}

type RegisterAuctionImage struct {
	ID                int32  `json:"id"`
	RegisterAuctionID int32  `json:"register_auction_id"`
	Url               string `json:"url"`
}

type Role struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID                  int32          `json:"id"`
	UserName            string         `json:"user_name"`
	Password            string         `json:"password"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	Address             string         `json:"address"`
	Phone               string         `json:"phone"`
	Birthdate           sql.NullTime   `json:"birthdate"`
	IDCard              string         `json:"id_card"`
	IDCardAddress       string         `json:"id_card_address"`
	IDCardDate          time.Time      `json:"id_card_date"`
	BankID              string         `json:"bank_id"`
	BankOwner           string         `json:"bank_owner"`
	BankName            string         `json:"bank_name"`
	Status              int32          `json:"status"`
	OrganizationName    sql.NullString `json:"organization_name"`
	OrganizationID      sql.NullString `json:"organization_id"`
	OrganizationDate    sql.NullTime   `json:"organization_date"`
	OrganizationAddress sql.NullString `json:"organization_address"`
	Position            sql.NullString `json:"position"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           sql.NullTime   `json:"updated_at"`
}

type UserImage struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"user_id"`
	Url    string `json:"url"`
	Type   int32  `json:"type"`
}

type UserRole struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}

// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}
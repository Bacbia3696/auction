// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Role struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID            int32        `json:"id"`
	Username      string       `json:"username"`
	Password      string       `json:"password"`
	Fullname      string       `json:"fullname"`
	Email         string       `json:"email"`
	Address       string       `json:"address"`
	Phone         string       `json:"phone"`
	Birthdate     sql.NullTime `json:"birthdate"`
	Idcard        string       `json:"idcard"`
	Idcardaddress string       `json:"idcardaddress"`
	Idcarddate    time.Time    `json:"idcarddate"`
	Bankid        string       `json:"bankid"`
	Bankowner     string       `json:"bankowner"`
	Bankname      string       `json:"bankname"`
	Status        int32        `json:"status"`
	Createdat     time.Time    `json:"createdat"`
	Updatedat     sql.NullTime `json:"updatedat"`
}

type UserImage struct {
	ID     int32  `json:"id"`
	Userid int32  `json:"userid"`
	Url    string `json:"url"`
}

type UserRole struct {
	ID     int32 `json:"id"`
	Userid int32 `json:"userid"`
	Roleid int32 `json:"roleid"`
}

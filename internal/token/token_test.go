package token

import (
	db "github.com/bacbia3696/auction/db/sqlc"
	"testing"
)

func TestBasic(t *testing.T) {
	user := db.User{
		UserName: "huongnq",
		ID:       1,
	}
	tokenString, err := GenToken(user)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := Verify(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Name != user.UserName {
		t.Errorf("Subject is not match, expected %s got %s", user.UserName, claims.Name)
	}
}

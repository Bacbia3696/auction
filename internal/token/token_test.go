package token

import (
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	username := "user1"
	userId :=1
	tokenString, err := GenToken(username, int32(userId), 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := Verify(tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if claims.Subject != username {
		t.Errorf("Subject is not match, expected %s got %s", username, claims.Subject)
	}
}

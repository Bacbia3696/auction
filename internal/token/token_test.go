package token

import (
	"testing"
)

func TestBasic(t *testing.T) {
	username := "huongnq"
	//userId :=1
	//tokenString, err := GenToken(nil, 5*time.Second)
	//if err != nil {
	//	t.Fatal(err)
	//}
	claims, err := Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiaHVvbmducSIsIklkIjo1LCJleHAiOjE2MzM2NTQ3MzN9.o0-6YktNafleNDxnvmJ_Hhs5hD7p1mAsIWmQwx_dbd0")
	if err != nil {
		t.Fatal(err)
	}
	if claims.Name != username {
		t.Errorf("Subject is not match, expected %s got %s", username, claims.Name)
	}
}

package paycode

import (
	"fmt"
	"testing"
)

func TestPaycodeMatch(t *testing.T) {
	tests := []struct {
		username string
		id       int
	}{
		{username: "username1", id: 100},
		{username: "username1", id: 123},
		{username: "xxx", id: 100},
		{username: "xxx", id: 123},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("username:%s_id:%d", test.username, test.id)
		t.Run(testname, func(t *testing.T) {
			code := NewFromAuctionID(test.username, test.id)
			rs := GetAuctionID(test.username, code)
			if test.id != rs {
				t.Errorf("id not mached, expected %d, got %d", test.id, rs)
			}
		})
	}
}

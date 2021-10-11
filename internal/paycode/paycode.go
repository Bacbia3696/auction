package paycode

import (
	"crypto/sha256"
	"strconv"

	"github.com/sirupsen/logrus"
)

func NewFromAuctionID(username string, auctionID int) int {
	key := sha256.Sum256([]byte(username))
	tp := []byte(strconv.Itoa(auctionID))
	rs := make([]byte, len(tp))
	for i := 0; i < len(tp); i++ {
		rs[i] = (tp[i]-'0'+key[i]%10)%10 + '0'
	}
	num, err := strconv.Atoi(string(rs))
	if err != nil {
		logrus.Info("err convert to number when gen paycode: ", err)
	}
	return num
}

func GetAuctionID(username string, paycode int) int {
	key := sha256.Sum256([]byte(username))
	tp := []byte(strconv.Itoa(paycode))
	rs := make([]byte, len(tp))
	for i := 0; i < len(tp); i++ {
		rs[i] = (tp[i]-'0'+10-key[i]%10)%10 + '0'
	}
	num, err := strconv.Atoi(string(rs))
	if err != nil {
		logrus.Info("err convert to number when gen paycode: ", err)
	}
	return num
}

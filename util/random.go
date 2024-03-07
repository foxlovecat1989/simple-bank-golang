package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqustuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(8)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "TWD", "RMB"}

	return currencies[rand.Intn(len(currencies))]
}

package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RamdomInt generates a ramdom interger between min and max
func RamdomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0 -> max - min + min
}

//RamdomString generates a ramdom string of length n
func RamdomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RamdomOwner generates a ramdom owner name
func RamdomOwner() string {
	return RamdomString(6)
}

//RamdomMoney generates a ramdom amount of money
func RamdomMoney() int64 {
	return RamdomInt(0, 1000000)
}

//RamdomString generates a ramdom string of length n
func RamdomCurrency() string {
	currencies := []string{"EUR", "USD", "CNY", "HKD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

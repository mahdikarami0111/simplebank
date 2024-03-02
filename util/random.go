package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int63n(max-min+1)
}

func RandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

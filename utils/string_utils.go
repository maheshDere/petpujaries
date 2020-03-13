package utils

import (
	"math/rand"
	"time"
)

const (
	passwordCharset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	passwordLength = 8
)

var (
	seededRand *rand.Rand
)

func init() {
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString() string {
	return StringWithCharset(passwordCharset, passwordLength)
}

func StringWithCharset(charSet string, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charSet[seededRand.Intn(len(charSet))]
	}
	return string(b)
}

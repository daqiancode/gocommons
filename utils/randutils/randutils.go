package randutils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	alphabetNumber       = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alphabetNumberNoCase = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	numbers              = []byte("0123456789")
	hex                  = []byte("abcdef0123456789")
)

func RandomNumber(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

func RandomAlphabetNumber(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabetNumber[rand.Intn(len(alphabetNumber))]
	}
	return string(b)
}

func RandomAlphabetNumberLower(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabetNumberNoCase[rand.Intn(len(alphabetNumberNoCase))]
	}
	return string(b)
}

func RandomHex(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = hex[rand.Intn(len(hex))]
	}
	return string(b)
}

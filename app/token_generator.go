package app

import (
	"crypto/rand"
	"math/big"
)

func GenerateToken(length int, symbols string) (token string) {
	token = ""
	var i int
	for len(token) < length {
		i = randomInt(len(symbols))
		token += string(symbols[i])
	}
	return
}

func randomInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}

package app

import (
	"crypto/rand"
	"math/big"
)

type dice struct {

}

func newDice() *dice {
	return &dice{}
}

func (d *dice) RollTheDice() []int {
	d1 := randomInt(1, 6)
	d2 := randomInt(1, 6)
	if d1 == d2 {
		return []int{d1, d2, d1, d2}
	}
	return []int{d1, d2}
}

func (d *dice) RollOneDice() int {
	return randomInt(1, 6)
}

func randomInt(min, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max - min + 1)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return min + int(n)
}
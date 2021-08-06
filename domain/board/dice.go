package board

import (
	"crypto/rand"
	"math/big"
)

type DiceState struct {
	Dice1 int
	Dice2 int
}

func (d *DiceState) IsEqualTo(d2 *DiceState) bool {
	if (d.Dice1==d2.Dice1 && d.Dice2==d2.Dice2) ||
		 (d.Dice1==d2.Dice2 && d.Dice2==d2.Dice1) {
		return  true
	}
	return false
}

func RollDice() *DiceState {
	return &DiceState{
		Dice1: RollOneDice(),
		Dice2: RollOneDice(),
	}
}

func RollOneDice() int {
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
//func (d *dice) RollTheDice() []int {
//	d1 := randomInt(1, 6)
//	d2 := randomInt(1, 6)
//	if d1 == d2 {
//		return []int{d1, d2, d1, d2}
//	}
//	return []int{d1, d2}
//
//
//}



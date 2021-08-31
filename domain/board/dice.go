package board

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type DiceState struct {
	Dice1 int
	Dice2 int
}

var ErrorDiceRandomGeneratorFailure = fmt.Errorf("dice random generator failure")

func (d *DiceState) IsEqualTo(d2 *DiceState) bool {
	if (d.Dice1 == d2.Dice1 && d.Dice2 == d2.Dice2) ||
		(d.Dice1 == d2.Dice2 && d.Dice2 == d2.Dice1) {
		return true
	}
	return false
}

func RollDice() (*DiceState, error) {
	dice1, err := RollOneDice()
	if err != nil {
		return &DiceState{}, nil
	}
	dice2, err := RollOneDice()
	if err != nil {
		return &DiceState{}, nil
	}

	return &DiceState{
		Dice1: dice1,
		Dice2: dice2,
	}, nil
}

func RollOneDice() (int, error) {
	return randomInt(1, 6)
}

func randomInt(min, max int) (int, error) {
	count := 0
	generated := false
	var n int
	for !generated {
		count++
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
		if err == nil {
			generated = true
			n = int(nBig.Int64())
		}
		if err != nil {
			time.Sleep(100 * time.Millisecond)
		}
		if count > 8 {
			return 0, ErrorDiceRandomGeneratorFailure
		}
	}
	return min + n, nil
}

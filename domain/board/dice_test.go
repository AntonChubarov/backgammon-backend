package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RollOneDice(t *testing.T) {
	var count [6]int
	n := 100000

	for i := 0; i < n; i++ {
		d:=RollOneDice()

		assert.GreaterOrEqual(t, d, 1)
		assert.LessOrEqual(t, d, 6)

		switch d {
		case 1:
			count[0]++
		case 2:
			count[1]++
		case 3:
			count[2]++
		case 4:
			count[3]++
		case 5:
			count[4]++
		case 6:
			count[5]++
		default:
			return
		}
	}

	for j := range count {
		assert.GreaterOrEqual(t, count[j], 3*n/6/4)
		assert.LessOrEqual(t, count[j], 5*n/6/4)
	}
}

func Test_RollDice(t *testing.T) {
	n := 100000
	var count [21]int
	diceCases := [21] DiceState{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{2, 2},
		{2, 3},
		{2, 4},
		{2, 5},
		{2, 6},
		{3, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{4, 4},
		{4, 5},
		{4, 6},
		{5, 5},
		{5, 6},
		{6, 6},
	}

	for i := 0; i < n; i++ {
		d := RollDice()

		assert.True(t, isDiceCaseValid(d))

		for j := range diceCases {
			if d.IsEqualTo(&diceCases[j]) {
				count[j]++
			}
		}
	}

	for k := range count {
		assert.Greater(t, count[k], 0)
	}
}

func isDiceCaseValid(d *DiceState) bool {
	if d.Dice1<1 || d.Dice1>6 || d.Dice2<1 || d.Dice2>6 { return false}
	return  true
}
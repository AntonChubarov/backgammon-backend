package game

import (
	"backgammon/domain/board"
	"backgammon/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DiceInterpreterTestCase struct {
	board.DiceState
	interpretation []int
}

func TestDiceInterpretationLongBackgammon (t *testing.T) {
	cases:=[]DiceInterpreterTestCase{
		{
			DiceState:      board.DiceState{
				Dice1: 1,
				Dice2: 2,
			},
			interpretation: []int{1,2},
		},
		{
			DiceState:      board.DiceState{
				Dice1: 6,
				Dice2: 4,
			},
			interpretation: []int{6,4},
		},
		{
			DiceState:      board.DiceState{
				Dice1: 3,
				Dice2: 3,
			},
			interpretation: []int{3,3,3,3},
		},
	}
	for i:=range cases {
		expected:= cases[i].interpretation
		actual:= DiceInterpretationLongBackgammon(&cases[i].DiceState)
		assert.Equal(t, true, utils.AreEqualIntSlices(expected, actual))
	}
}

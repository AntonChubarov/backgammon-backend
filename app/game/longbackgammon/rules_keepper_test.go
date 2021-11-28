package longbackgammon

import (
	"backgammon/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var longBackgammonRulesKeepper = NewLongBackgammonRulesKeepper()

type completeTurnTestCase struct {
	Game          *Game
	Color         Color
	Turn          *Turn
	ConsumedDice  []int
	ExpectedError error
}

func TestLongBackgammonRulesKeepper(t *testing.T) {
	cases := []completeTurnTestCase{
		{
			Game: &Game{
				State:              InProcess,
				AwaitingTurnNumber: 1,
				CurrentTurn:        Black,
			},
			Color:         Black,
			Turn:          &Turn{TurnNumber: 1},
			ConsumedDice:  nil,
			ExpectedError: nil,
		},
	}

	for i := range cases {
		expected := cases[i].ExpectedError
		actual := longBackgammonRulesKeepper.ValidateAllRules(cases[i].Game, cases[i].Color, cases[i].Turn, cases[i].ConsumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestDiceInterpretationLongBackgammon(t *testing.T) {
	cases := []DiceInterpreterTestCase{
		{
			DiceState: DiceState{
				Dice1: 1,
				Dice2: 2,
			},
			interpretation: []int{1, 2},
		},
		{
			DiceState: DiceState{
				Dice1: 6,
				Dice2: 4,
			},
			interpretation: []int{6, 4},
		},
		{
			DiceState: DiceState{
				Dice1: 3,
				Dice2: 3,
			},
			interpretation: []int{3, 3, 3, 3},
		},
	}
	for i := range cases {
		expected := cases[i].interpretation
		actual := DiceInterpretationLongBackgammon(&cases[i].DiceState)
		assert.Equal(t, true, utils.AreEqualIntSlices(expected, actual))
	}
}

package game

import (
	"backgammon/domain/board"
	"backgammon/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var longBackgammonRulesKeepper = NewLongBackgammonRulesKeepper()

type completeTurnTestCase struct {
	Game          *Game
	Color         board.StickColor
	Turn          *board.Turn
	ConsumedDice  []int
	ExpectedError error
}

func TestLongBackgammonRulesKeepper(t *testing.T) {
	cases := []completeTurnTestCase{
		{
			Game: &Game{
				State:              InProcess,
				AwaitingTurnNumber: 1,
				CurrentTurn:        board.Black,
			},
			Color:         board.Black,
			Turn:          &board.Turn{TurnNumber: 1},
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
			DiceState: board.DiceState{
				Dice1: 1,
				Dice2: 2,
			},
			interpretation: []int{1, 2},
		},
		{
			DiceState: board.DiceState{
				Dice1: 6,
				Dice2: 4,
			},
			interpretation: []int{6, 4},
		},
		{
			DiceState: board.DiceState{
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

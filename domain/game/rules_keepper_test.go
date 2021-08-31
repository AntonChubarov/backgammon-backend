package game

import (
	"backgammon/domain/board"
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

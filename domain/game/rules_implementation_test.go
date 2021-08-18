package game

import (
	"backgammon/domain/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Move *board.Move
	ExpectedError error
}

func TestRuleMoveMatchStickColor (t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{From: 13, To: 15},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &board.Move{From: 14, To: 16},
			ExpectedError: nil,
		},
	}

	holes := [24]board.Hole{}

	holes[0] = board.Hole{
		StickColor: board.Black,
		StickCount: 15,
	}

	holes[12] = board.Hole{
		StickColor: board.White,
		StickCount: 15,
	}

	gameBoard := board.Board{Holes: holes}

	g := &Game{Board: gameBoard}
	c := board.White
	consumedDice := []int{2, 3}


	rule := RuleMoveMatchStickColor{nextRule: nil}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, c, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}
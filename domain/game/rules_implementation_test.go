package game

import (
	"backgammon/domain/board"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type testCase struct {
	Move *board.Move
	ExpectedError error
	color board.StickColor
}

func TestRuleMoveMatchStickColor (t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{From: 13, To: 15},
			color: board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			color: board.White,
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &board.Move{From: 14, To: 16},
			color: board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 13, To: 15},
			color: board.Black,
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			color: board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 14, To: 16},
			color: board.Black,
			ExpectedError: nil,
		},
	}

	//holes := [25]board.Hole{}
	//
	//holes[1] = board.Hole{
	//	StickColor: board.Black,
	//	StickCount: 15,
	//}
	//
	//holes[13] = board.Hole{
	//	StickColor: board.White,
	//	StickCount: 15,
	//}
	//
	//gameBoard := board.Board{Holes: holes}
	gameBoard:=board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor=board.Black
	gameBoard.Holes[1].StickCount=15
	gameBoard.Holes[13].StickColor=board.White
	gameBoard.Holes[13].StickCount=15




	g := &Game{Board: gameBoard}

	var consumedDice  []int


	rule := RuleMoveMatchStickColor{nextRule: nil}
	for i := range cases {
		log.Println("test case: ", i)
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

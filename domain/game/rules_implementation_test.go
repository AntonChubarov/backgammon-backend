package game

import (
	"backgammon/domain/board"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type gameRuleTestCase struct {
	Game          *Game
	Color         board.StickColor
	ExpectedError error
}

type testCase struct {
	Move          *board.Move
	ExpectedError error
	color         board.StickColor
}

func TestRuleMatchOrder(t *testing.T) {
	cases := []gameRuleTestCase{
		{
			Game: &Game{
				CurrentTurn: board.White,
			},
			Color:         board.White,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				CurrentTurn: board.Black,
			},
			Color:         board.Black,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				CurrentTurn: board.White,
			},
			Color:         board.Black,
			ExpectedError: ErrorOutOfTurn,
		},
		{
			Game: &Game{
				CurrentTurn: board.Black,
			},
			Color:         board.White,
			ExpectedError: ErrorOutOfTurn,
		},
	}

	rule := RuleMatchOrder{}

	for i := range cases {
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color))
	}
}

func TestRuleCorrectGamePhase_ValidateRule(t *testing.T) {
	cases := []gameRuleTestCase{
		{
			Game: &Game{
				State: InProcess,
			},
			Color:         board.Black,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				State: NotStarted,
			},
			Color:         board.Black,
			ExpectedError: ErrorUotOfGame,
		},
		{
			Game: &Game{
				State: Finished,
			},
			Color:         board.Black,
			ExpectedError: ErrorUotOfGame,
		},
		{
			Game: &Game{
				State: InProcess,
			},
			Color:         board.White,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				State: NotStarted,
			},
			Color:         board.White,
			ExpectedError: ErrorUotOfGame,
		},
		{
			Game: &Game{
				State: Finished,
			},
			Color:         board.White,
			ExpectedError: ErrorUotOfGame,
		},
	}

	rule := RuleCorrectGamePhase{}

	for i := range cases {
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color))
	}
}

func TestRuleMoveMatchStickColor(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{From: 13, To: 15},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			color:         board.White,
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &board.Move{From: 14, To: 16},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 13, To: 15},
			color:         board.Black,
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			color:         board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 14, To: 16},
			color:         board.Black,
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
	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleMoveMatchStickColor{nextRule: nil}
	for i := range cases {
		log.Println("test case: ", i)
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

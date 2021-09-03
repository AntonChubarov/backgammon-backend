package game

import (
	"backgammon/domain/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

type gameRuleTestCase struct {
	Game          *Game
	Color         board.StickColor
	Turn          board.Turn
	ExpectedError error
}

type testCase struct {
	Move          *board.Move
	ExpectedError error
	color         board.StickColor
}

type turnRuleTestCase struct {
	Color         board.StickColor
	Turn          *board.Turn
	ExpectedError error
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
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color, &board.Turn{}))
	}
}

func TestRuleCorrectGamePhase(t *testing.T) {
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
			ExpectedError: ErrorOutOfGame,
		},
		{
			Game: &Game{
				State: Finished,
			},
			Color:         board.Black,
			ExpectedError: ErrorOutOfGame,
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
			ExpectedError: ErrorOutOfGame,
		},
		{
			Game: &Game{
				State: Finished,
			},
			Color:         board.White,
			ExpectedError: ErrorOutOfGame,
		},
	}

	rule := RuleCorrectGamePhase{}

	for i := range cases {
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color, &board.Turn{}))
	}
}

func TestRuleMatchTurnNumber(t *testing.T) {
	cases := []gameRuleTestCase{
		{
			Game: &Game{
				AwaitingTurnNumber: 1,
			},
			Turn:          board.Turn{TurnNumber: 1},
			ExpectedError: nil,
		},
		{
			Game: &Game{
				AwaitingTurnNumber: 1,
			},
			Turn:          board.Turn{TurnNumber: 0},
			ExpectedError: ErrorInvalidTurnNumber,
		},
		{
			Game: &Game{
				AwaitingTurnNumber: 1,
			},
			Turn:          board.Turn{TurnNumber: 2},
			ExpectedError: ErrorInvalidTurnNumber,
		},
	}

	rule := RuleMatchTurnNumber{}

	for i := range cases {
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color, &cases[i].Turn))
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
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleMoveDirection(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{From: 13, To: 15},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 16, To: 14},
			color:         board.White,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
		{
			Move:          &board.Move{From: 16, To: 16},
			color:         board.White,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
		{
			Move:          &board.Move{From: 13, To: 15},
			color:         board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 1, To: 3},
			color:         board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 16, To: 14},
			color:         board.Black,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
		{
			Move:          &board.Move{From: 16, To: 16},
			color:         board.Black,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleMoveDirection{nextRule: nil}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleMoveToOccupiedHole_ValidateRule(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{From: 23, To: 2},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 22, To: 3},
			color:         board.White,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
		{
			Move:          &board.Move{From: 21, To: 1},
			color:         board.White,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
		{
			Move:          &board.Move{From: 12, To: 14},
			color:         board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 13, To: 15},
			color:         board.Black,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
		{
			Move:          &board.Move{From: 11, To: 13},
			color:         board.Black,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 5
	for i := 3; i <= 12; i++ {
		gameBoard.Holes[i].StickColor = board.Black
		gameBoard.Holes[i].StickCount = 1
	}
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 5
	for i := 15; i <= 24; i++ {
		gameBoard.Holes[i].StickColor = board.White
		gameBoard.Holes[i].StickCount = 1
	}

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleMoveToOccupiedHole{nextRule: nil}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleMoveFromEmptyHole(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{From: 13, To: 16},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 14, To: 16},
			color:         board.White,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 16, To: 18},
			color:         board.White,
			ExpectedError: ErrorMoveFromEmptyHole,
		},
		{
			Move:          &board.Move{From: 1, To: 4},
			color:         board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 2, To: 4},
			color:         board.Black,
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{From: 4, To: 6},
			color:         board.Black,
			ExpectedError: ErrorMoveFromEmptyHole,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 13
	for i := 2; i <= 3; i++ {
		gameBoard.Holes[i].StickColor = board.Black
		gameBoard.Holes[i].StickCount = 1
	}
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 13
	for i := 14; i <= 15; i++ {
		gameBoard.Holes[i].StickColor = board.White
		gameBoard.Holes[i].StickCount = 1
	}

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleMoveFromEmptyHole{nextRule: nil}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleForbiddenMoveKindLongBackgammon(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{MoveKind: board.Movement},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: board.Surrender},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: 3},
			ExpectedError: ErrorImpossibleMoveKind,
		},
		{
			Move:          &board.Move{MoveKind: 998},
			ExpectedError: ErrorImpossibleMoveKind,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleForbiddenMoveKindLongBackgammon{nextRule: nil}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleMoveFormat(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{MoveKind: board.Movement, From: 1, To: 3},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: board.Movement, From: 0, To: 3},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &board.Move{MoveKind: board.Movement, From: 24, To: 1},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: board.Movement, From: 23, To: 25},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &board.Move{MoveKind: board.Movement, From: 24, To: 0},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 24, To: 1},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 24, To: 0},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: board.Surrender, From: 1, To: 3},
			ExpectedError: nil,
		},
		{
			Move:          &board.Move{MoveKind: 5, From: 1, To: 3},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleMoveFormat{}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleRemovingNotFromHome(t *testing.T) {
	cases := []testCase{
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 24, To: 0},
			ExpectedError: nil,
			color:         board.Black,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 19, To: 0},
			ExpectedError: nil,
			color:         board.Black,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 18, To: 0},
			ExpectedError: ErrorRemovingFromInvalidHole,
			color:         board.Black,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 12, To: 0},
			ExpectedError: nil,
			color:         board.White,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 7, To: 0},
			ExpectedError: nil,
			color:         board.White,
		},
		{
			Move:          &board.Move{MoveKind: board.Removing, From: 6, To: 0},
			ExpectedError: ErrorRemovingFromInvalidHole,
			color:         board.White,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	for i := 19; i <= 24; i++ {
		gameBoard.Holes[i].StickColor = board.Black
		gameBoard.Holes[i].StickCount = 2
	}

	for i := 7; i <= 12; i++ {
		gameBoard.Holes[i].StickColor = board.White
		gameBoard.Holes[i].StickCount = 2
	}

	g := &Game{Board: gameBoard}

	var consumedDice []int

	rule := RuleRemovingNotFromHome{}
	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].color, cases[i].Move, consumedDice)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleTooMuchSteps_NonPairDice(t *testing.T) {
	cases := []turnRuleTestCase{
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: board.Movement,
						From:     2,
						To:       4,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: board.Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: board.Movement,
						From:     4,
						To:       7,
					},
				},
			},
			ExpectedError: ErrorTooMuchStepsInTurn,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{
		Board: gameBoard,
		DiceState: board.DiceState{
			Dice1: 1,
			Dice2: 2,
		},
	}

	rule := RuleTooMuchSteps{}

	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].Color, cases[i].Turn)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleTooMuchSteps_PairDice(t *testing.T) {
	cases := []turnRuleTestCase{
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: board.Movement,
						From:     2,
						To:       4,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: board.Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: board.Movement,
						From:     4,
						To:       7,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: board.Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: board.Movement,
						From:     4,
						To:       7,
					},
					{
						MoveKind: board.Movement,
						From:     7,
						To:       11,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &board.Turn{
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: board.Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: board.Movement,
						From:     4,
						To:       7,
					},
					{
						MoveKind: board.Movement,
						From:     7,
						To:       11,
					},
					{
						MoveKind: board.Movement,
						From:     7,
						To:       16,
					},
				},
			},
			ExpectedError: ErrorTooMuchStepsInTurn,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{
		Board: gameBoard,
		DiceState: board.DiceState{
			Dice1: 1,
			Dice2: 1,
		},
	}

	rule := RuleTooMuchSteps{}

	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].Color, cases[i].Turn)
		assert.Equal(t, expected, actual)
	}
}

func TestRuleAttemptToGetFewSticksFromHead(t *testing.T) {
	cases := []turnRuleTestCase{
		{
			Color: board.Black,
			Turn: &board.Turn{
				StickColor: board.Black,
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       3,
					},
					{
						MoveKind: board.Movement,
						From:     3,
						To:       6,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Color: board.Black,
			Turn: &board.Turn{
				StickColor: board.Black,
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       3,
					},
					{
						MoveKind: board.Movement,
						From:     1,
						To:       4,
					},
				},
			},
			ExpectedError: ErrorMoveFromHeadLimit1,
		},
		{
			Color: board.Black,
			Turn: &board.Turn{
				StickColor: board.Black,
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     13,
						To:       15,
					},
					{
						MoveKind: board.Movement,
						From:     13,
						To:       16,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Color: board.White,
			Turn: &board.Turn{
				StickColor: board.White,
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     13,
						To:       15,
					},
					{
						MoveKind: board.Movement,
						From:     15,
						To:       18,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Color: board.White,
			Turn: &board.Turn{
				StickColor: board.White,
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     13,
						To:       15,
					},
					{
						MoveKind: board.Movement,
						From:     13,
						To:       16,
					},
				},
			},
			ExpectedError: ErrorMoveFromHeadLimit1,
		},
		{
			Color: board.White,
			Turn: &board.Turn{
				StickColor: board.White,
				Moves: []board.Move{
					{
						MoveKind: board.Movement,
						From:     1,
						To:       3,
					},
					{
						MoveKind: board.Movement,
						From:     1,
						To:       4,
					},
				},
			},
			ExpectedError: nil,
		},
	}

	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{Board: gameBoard}

	rule := RuleAttemptToGetFewSticksFromHead{}

	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].Color, cases[i].Turn)
		assert.Equal(t, expected, actual)
	}
}

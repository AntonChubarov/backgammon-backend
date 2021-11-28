package longbackgammon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type gameRuleTestCase struct {
	Game          *Game
	Color         Color
	Turn          Turn
	ExpectedError error
}

type testCase struct {
	Move          *Move
	ExpectedError error
	color         Color
}

type turnRuleTestCase struct {
	Color         Color
	Turn          *Turn
	ExpectedError error
}

func TestRuleMatchOrder(t *testing.T) {
	cases := []gameRuleTestCase{
		{
			Game: &Game{
				CurrentTurn: White,
			},
			Color:         White,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				CurrentTurn: Black,
			},
			Color:         Black,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				CurrentTurn: White,
			},
			Color:         Black,
			ExpectedError: ErrorOutOfTurn,
		},
		{
			Game: &Game{
				CurrentTurn: Black,
			},
			Color:         White,
			ExpectedError: ErrorOutOfTurn,
		},
	}

	rule := RuleMatchOrder{}

	for i := range cases {
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color, &Turn{}))
	}
}

func TestRuleCorrectGamePhase(t *testing.T) {
	cases := []gameRuleTestCase{
		{
			Game: &Game{
				State: InProcess,
			},
			Color:         Black,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				State: NotStarted,
			},
			Color:         Black,
			ExpectedError: ErrorOutOfGame,
		},
		{
			Game: &Game{
				State: Finished,
			},
			Color:         Black,
			ExpectedError: ErrorOutOfGame,
		},
		{
			Game: &Game{
				State: InProcess,
			},
			Color:         White,
			ExpectedError: nil,
		},
		{
			Game: &Game{
				State: NotStarted,
			},
			Color:         White,
			ExpectedError: ErrorOutOfGame,
		},
		{
			Game: &Game{
				State: Finished,
			},
			Color:         White,
			ExpectedError: ErrorOutOfGame,
		},
	}

	rule := RuleCorrectGamePhase{}

	for i := range cases {
		assert.Equal(t, cases[i].ExpectedError, rule.ValidateRule(cases[i].Game, cases[i].Color, &Turn{}))
	}
}

func TestRuleMatchTurnNumber(t *testing.T) {
	cases := []gameRuleTestCase{
		{
			Game: &Game{
				AwaitingTurnNumber: 1,
			},
			Turn:          Turn{TurnNumber: 1},
			ExpectedError: nil,
		},
		{
			Game: &Game{
				AwaitingTurnNumber: 1,
			},
			Turn:          Turn{TurnNumber: 0},
			ExpectedError: ErrorInvalidTurnNumber,
		},
		{
			Game: &Game{
				AwaitingTurnNumber: 1,
			},
			Turn:          Turn{TurnNumber: 2},
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
			Move:          &Move{From: 13, To: 15},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 1, To: 3},
			color:         White,
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &Move{From: 14, To: 16},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 13, To: 15},
			color:         Black,
			ExpectedError: ErrorOpponentsStickMoveAttempt,
		},
		{
			Move:          &Move{From: 1, To: 3},
			color:         Black,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 14, To: 16},
			color:         Black,
			ExpectedError: nil,
		},
	}

	//holes := [25]board.Hole{}
	//
	//holes[1] = board.Hole{
	//	Color: board.Black,
	//	StickCount: 15,
	//}
	//
	//holes[13] = board.Hole{
	//	Color: board.White,
	//	StickCount: 15,
	//}
	//
	//gameBoard := board.Board{Holes: holes}
	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].Color = White
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
			Move:          &Move{From: 13, To: 15},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 1, To: 3},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 16, To: 14},
			color:         White,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
		{
			Move:          &Move{From: 16, To: 16},
			color:         White,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
		{
			Move:          &Move{From: 13, To: 15},
			color:         Black,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 1, To: 3},
			color:         Black,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 16, To: 14},
			color:         Black,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
		{
			Move:          &Move{From: 16, To: 16},
			color:         Black,
			ExpectedError: ErrorIncorrectMoveDirection,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].Color = White
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
			Move:          &Move{From: 23, To: 2},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 22, To: 3},
			color:         White,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
		{
			Move:          &Move{From: 21, To: 1},
			color:         White,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
		{
			Move:          &Move{From: 12, To: 14},
			color:         Black,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 13, To: 15},
			color:         Black,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
		{
			Move:          &Move{From: 11, To: 13},
			color:         Black,
			ExpectedError: ErrorMoveToOccupiedHole,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 5
	for i := 3; i <= 12; i++ {
		gameBoard.Holes[i].Color = Black
		gameBoard.Holes[i].StickCount = 1
	}
	gameBoard.Holes[13].Color = White
	gameBoard.Holes[13].StickCount = 5
	for i := 15; i <= 24; i++ {
		gameBoard.Holes[i].Color = White
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
			Move:          &Move{From: 13, To: 16},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 14, To: 16},
			color:         White,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 16, To: 18},
			color:         White,
			ExpectedError: ErrorMoveFromEmptyHole,
		},
		{
			Move:          &Move{From: 1, To: 4},
			color:         Black,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 2, To: 4},
			color:         Black,
			ExpectedError: nil,
		},
		{
			Move:          &Move{From: 4, To: 6},
			color:         Black,
			ExpectedError: ErrorMoveFromEmptyHole,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 13
	for i := 2; i <= 3; i++ {
		gameBoard.Holes[i].Color = Black
		gameBoard.Holes[i].StickCount = 1
	}
	gameBoard.Holes[13].Color = White
	gameBoard.Holes[13].StickCount = 13
	for i := 14; i <= 15; i++ {
		gameBoard.Holes[i].Color = White
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
			Move:          &Move{MoveKind: Movement},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: Removing},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: Surrender},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: 3},
			ExpectedError: ErrorImpossibleMoveKind,
		},
		{
			Move:          &Move{MoveKind: 998},
			ExpectedError: ErrorImpossibleMoveKind,
		},
	}

	gameBoard := Board{}
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
			Move:          &Move{MoveKind: Movement, From: 1, To: 3},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: Movement, From: 0, To: 3},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &Move{MoveKind: Movement, From: 24, To: 1},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: Movement, From: 23, To: 25},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &Move{MoveKind: Movement, From: 24, To: 0},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 24, To: 1},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 24, To: 0},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: Surrender, From: 1, To: 3},
			ExpectedError: nil,
		},
		{
			Move:          &Move{MoveKind: 5, From: 1, To: 3},
			ExpectedError: ErrorIncorrectMoveFormat,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].Color = White
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
			Move:          &Move{MoveKind: Removing, From: 24, To: 0},
			ExpectedError: nil,
			color:         Black,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 19, To: 0},
			ExpectedError: nil,
			color:         Black,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 18, To: 0},
			ExpectedError: ErrorRemovingFromInvalidHole,
			color:         Black,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 12, To: 0},
			ExpectedError: nil,
			color:         White,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 7, To: 0},
			ExpectedError: nil,
			color:         White,
		},
		{
			Move:          &Move{MoveKind: Removing, From: 6, To: 0},
			ExpectedError: ErrorRemovingFromInvalidHole,
			color:         White,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	for i := 19; i <= 24; i++ {
		gameBoard.Holes[i].Color = Black
		gameBoard.Holes[i].StickCount = 2
	}

	for i := 7; i <= 12; i++ {
		gameBoard.Holes[i].Color = White
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
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: Movement,
						From:     2,
						To:       4,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: Movement,
						From:     4,
						To:       7,
					},
				},
			},
			ExpectedError: ErrorTooMuchStepsInTurn,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].Color = White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{
		Board: gameBoard,
		DiceState: DiceState{
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
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: Movement,
						From:     2,
						To:       4,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: Movement,
						From:     4,
						To:       7,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: Movement,
						From:     4,
						To:       7,
					},
					{
						MoveKind: Movement,
						From:     7,
						To:       11,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Turn: &Turn{
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       2,
					},
					{
						MoveKind: Movement,
						From:     2,
						To:       4,
					},
					{
						MoveKind: Movement,
						From:     4,
						To:       7,
					},
					{
						MoveKind: Movement,
						From:     7,
						To:       11,
					},
					{
						MoveKind: Movement,
						From:     7,
						To:       16,
					},
				},
			},
			ExpectedError: ErrorTooMuchStepsInTurn,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].Color = White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{
		Board: gameBoard,
		DiceState: DiceState{
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
			Color: Black,
			Turn: &Turn{
				Color: Black,
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       3,
					},
					{
						MoveKind: Movement,
						From:     3,
						To:       6,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Color: Black,
			Turn: &Turn{
				Color: Black,
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       3,
					},
					{
						MoveKind: Movement,
						From:     1,
						To:       4,
					},
				},
			},
			ExpectedError: ErrorMoveFromHeadLimit1,
		},
		{
			Color: Black,
			Turn: &Turn{
				Color: Black,
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     13,
						To:       15,
					},
					{
						MoveKind: Movement,
						From:     13,
						To:       16,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Color: White,
			Turn: &Turn{
				Color: White,
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     13,
						To:       15,
					},
					{
						MoveKind: Movement,
						From:     15,
						To:       18,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Color: White,
			Turn: &Turn{
				Color: White,
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     13,
						To:       15,
					},
					{
						MoveKind: Movement,
						From:     13,
						To:       16,
					},
				},
			},
			ExpectedError: ErrorMoveFromHeadLimit1,
		},
		{
			Color: White,
			Turn: &Turn{
				Color: White,
				Moves: []Move{
					{
						MoveKind: Movement,
						From:     1,
						To:       3,
					},
					{
						MoveKind: Movement,
						From:     1,
						To:       4,
					},
				},
			},
			ExpectedError: nil,
		},
	}

	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 15
	gameBoard.Holes[13].Color = White
	gameBoard.Holes[13].StickCount = 15

	g := &Game{Board: gameBoard}

	rule := RuleAttemptToGetFewSticksFromHead{}

	for i := range cases {
		expected := cases[i].ExpectedError
		actual := rule.ValidateRule(g, cases[i].Color, cases[i].Turn)
		assert.Equal(t, expected, actual)
	}
}

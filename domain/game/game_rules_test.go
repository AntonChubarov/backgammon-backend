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

type StartOfFenceCase struct {
	Board board.Board
	Color board.StickColor
	StartHole int
	IsStartOfFence bool
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

func TestIsStartOfFence(t *testing.T) {
	gameBoard:=board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor=board.Black
	gameBoard.Holes[1].StickCount=9
	for i := 5; i <= 10; i++ {
		gameBoard.Holes[i].StickColor=board.Black
		gameBoard.Holes[i].StickCount=1
	}
	gameBoard.Holes[13].StickColor=board.White
	gameBoard.Holes[13].StickCount=15


	cases:=[]StartOfFenceCase{
		{
			Board:          gameBoard,
			Color: board.Black,
			StartHole:      5,
			IsStartOfFence: true,
		},
		{
			Board:          gameBoard,
			Color: board.White,
			StartHole:      5,
			IsStartOfFence: false,
		},
		{
			Board:          gameBoard,
			Color: board.Black,
			StartHole:      4,
			IsStartOfFence: false,
		},
	}

	for i := range cases {
		actual := IsStartOfFence(cases[i].Board, cases[i].Color, cases[i].StartHole)
		expected := cases[i].IsStartOfFence
		assert.Equal(t, expected, actual)
	}
}

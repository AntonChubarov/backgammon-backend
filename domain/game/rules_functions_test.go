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
	Board          board.Board
	Color          board.StickColor
	StartHole      int
	IsStartOfFence bool
}


func TestIsStartOfFence(t *testing.T) {
	gameBoard := board.Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].StickColor = board.Black
	gameBoard.Holes[1].StickCount = 9
	for i := 5; i <= 10; i++ {
		gameBoard.Holes[i].StickColor = board.Black
		gameBoard.Holes[i].StickCount = 1
	}
	gameBoard.Holes[13].StickColor = board.White
	gameBoard.Holes[13].StickCount = 9
	for i := 17; i <= 22; i++ {
		gameBoard.Holes[i].StickColor = board.White
		gameBoard.Holes[i].StickCount = 1
	}

	cases := []StartOfFenceCase{
		{
			Board:          gameBoard,
			Color:          board.Black,
			StartHole:      5,
			IsStartOfFence: true,
		},
		{
			Board:          gameBoard,
			Color:          board.White,
			StartHole:      5,
			IsStartOfFence: false,
		},
		{
			Board:          gameBoard,
			Color:          board.Black,
			StartHole:      4,
			IsStartOfFence: false,
		},
		{
			Board:          gameBoard,
			Color:          board.White,
			StartHole:      17,
			IsStartOfFence: true,
		},
	}

	for i := range cases {
		actual := IsStartOfFence(cases[i].Board, cases[i].Color, cases[i].StartHole)
		expected := cases[i].IsStartOfFence
		assert.Equal(t, expected, actual)
	}
}

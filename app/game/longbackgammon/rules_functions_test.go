package longbackgammon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DiceInterpreterTestCase struct {
	DiceState
	interpretation []int
}

type StartOfFenceCase struct {
	Board          Board
	Color          Color
	StartHole      int
	IsStartOfFence bool
}

func TestIsStartOfFence(t *testing.T) {
	gameBoard := Board{}
	gameBoard.Clear()
	gameBoard.Holes[1].Color = Black
	gameBoard.Holes[1].StickCount = 9
	for i := 5; i <= 10; i++ {
		gameBoard.Holes[i].Color = Black
		gameBoard.Holes[i].StickCount = 1
	}
	gameBoard.Holes[13].Color = White
	gameBoard.Holes[13].StickCount = 9
	for i := 17; i <= 22; i++ {
		gameBoard.Holes[i].Color = White
		gameBoard.Holes[i].StickCount = 1
	}

	cases := []StartOfFenceCase{
		{
			Board:          gameBoard,
			Color:          Black,
			StartHole:      5,
			IsStartOfFence: true,
		},
		{
			Board:          gameBoard,
			Color:          White,
			StartHole:      5,
			IsStartOfFence: false,
		},
		{
			Board:          gameBoard,
			Color:          Black,
			StartHole:      4,
			IsStartOfFence: false,
		},
		{
			Board:          gameBoard,
			Color:          White,
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

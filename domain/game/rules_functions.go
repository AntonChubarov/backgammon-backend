package game

import (
	"backgammon/domain/board"
)

func MoveDistance(c board.StickColor, from int, to int) int {
	if c==board.Black {
		return to-from
	} else {
		return board.InvertNumeration(to)-board.InvertNumeration(from)
	}
}

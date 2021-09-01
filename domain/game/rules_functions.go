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

func IsStartOfFence(b board.Board, c board.StickColor, startHole int) bool {
	if c == board.Black {
		for j := startHole; j <= startHole+5 && j <= 24; j++ {
			if b.Holes[j].StickColor != c {
				return false
			}
		}
		return true
	} else {
		startHole = board.InvertNumeration(startHole)
		for j := startHole; j <= startHole+5 && j <= 24; j++ {
			i := board.InvertNumeration(j)
			if b.Holes[i].StickColor != c {
				return false
			}
		}
		return true
	}
}

//func IsFenceBlocking(b board.Board, fenceStartHole int)
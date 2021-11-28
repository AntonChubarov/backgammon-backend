package longbackgammon

func MoveDistance(c Color, from int, to int) int {
	if c == Black {
		return to - from
	} else {
		return InvertNumeration(to) - InvertNumeration(from)
	}
}

func IsStartOfFence(b Board, c Color, startHole int) bool {
	if c == Black {
		for j := startHole; j <= startHole+5 && j <= 24; j++ {
			if b.Holes[j].Color != c {
				return false
			}
		}
		return true
	} else {
		startHole = InvertNumeration(startHole)
		for j := startHole; j <= startHole+5 && j <= 24; j++ {
			i := InvertNumeration(j)
			if b.Holes[i].Color != c {
				return false
			}
		}
		return true
	}
}

//func IsFenceBlocking(b board.Board, fenceStartHole int)

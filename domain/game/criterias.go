package game

import "backgammon/domain/board"

type Criteria interface {
	ValidateCriteria(b *board.Board,sc board.StickColor) error
}




package game

import "backgammon/domain/board"


type MovingRule interface {
	ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error
	SetNextRule(mr MovingRule)
}

type TurnRule interface {
	ValidateRule(g *Game, t *board.Turn) error
	SetNextRule(tr TurnRule)
}

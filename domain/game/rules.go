package game

import "backgammon/domain/board"

type GameRule interface {
	ValidateRule(g *Game, c board.StickColor, t *board.Turn) error
	SetNextRule(gr GameRule)
}

type MovingRule interface {
	ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error
	SetNextRule(mr MovingRule)
}

type TurnRule interface {
	ValidateRule(g *Game, c board.StickColor, t *board.Turn) error
	SetNextRule(tr TurnRule)
}

package longbackgammon

type MovingRule interface {
	ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error
	SetNextRule(mr MovingRule)
}

type TurnRule interface {
	ValidateRule(g *Game, t *Turn) error
	SetNextRule(tr TurnRule)
}

package game

import "backgammon/domain/board"

type LongBackgammonRulesKeepper struct {
	initialGameRule   GameRule
	initialMovingRule MovingRule
	initialTurnRule   TurnRule
}

func NewLongBackgammonRulesKeepper() *LongBackgammonRulesKeepper {
	gr1 := &RuleCorrectGamePhase{}
	gr2 := &RuleMatchTurnNumber{}
	gr1.SetNextRule(gr2)
	gr3 := &RuleMatchOrder{}
	gr2.SetNextRule(gr3)

	return &LongBackgammonRulesKeepper{initialGameRule: gr1}
}

func (lbrk *LongBackgammonRulesKeepper) ValidateAllRules(g *Game, c board.StickColor, t *board.Turn, consumedDice []int) (err error) {
	err = lbrk.initialGameRule.ValidateRule(g, c, t)
	if err != nil {
		return
	}

	return
}

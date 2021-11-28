package longbackgammon

type LongBackgammonRulesKeepper struct {
	initialGameRule   TurnRule
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

func (lbrk *LongBackgammonRulesKeepper) ValidateAllRules(g *Game, c Color, t *Turn, consumedDice []int) (err error) {
	err = lbrk.initialGameRule.ValidateRule(g, t)
	if err != nil {
		return
	}

	return
}

func DiceInterpretationLongBackgammon(d *DiceState) []int {
	var steps []int
	if d.Dice1 == d.Dice2 {
		steps = make([]int, 4, 4)
		steps[0] = d.Dice1
		steps[1] = d.Dice1
		steps[2] = d.Dice1
		steps[3] = d.Dice1
		return steps

	}
	steps = make([]int, 2, 2)
	steps[0] = d.Dice1
	steps[1] = d.Dice2
	return steps
}

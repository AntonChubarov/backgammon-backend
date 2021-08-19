package game

func (r *RuleMoveMatchStickColor) SetNextRule(mr MovingRule) {
	r.nextRule=mr
}

func (r *RuleMoveDirection) SetNextRule(mr MovingRule) {
	r.nextRule=mr
}

func (r *RuleMoveImpossibleAmountSteps) SetNextRule(mr MovingRule) {
	r.nextRule=mr
}
func (r *RuleMoveToOccupiedHole) SetNextRule(mr MovingRule) {
	r.nextRule=mr
}
func (r *RuleMoveFromEmptyHole) SetNextRule(mr MovingRule) {
	r.nextRule=mr
}

func (r *RuleForbiddenMoveKindInGameType) SetNextRule(mr MovingRule) {
	r.nextRule=mr
}




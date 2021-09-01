package game

func (r *RuleMatchOrder) SetNextRule(gr TurnRule) {
	r.nextRule = gr
}

func (r *RuleCorrectGamePhase) SetNextRule(gr TurnRule) {
	r.nextRule = gr
}

func (r *RuleMatchTurnNumber) SetNextRule(gr TurnRule) {
	r.nextRule = gr
}

func (r *RuleMoveMatchStickColor) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

func (r *RuleMoveDirection) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

func (r *RuleMoveImpossibleAmountSteps) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}
func (r *RuleMoveToOccupiedHole) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}
func (r *RuleMoveFromEmptyHole) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

func (r *RuleMoveFormat) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

func (r *RuleRemovingNotFromHome) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

// Need to check this type
//func (r *RuleForbiddenMoveKindInGameType) SetNextRule(mr MovingRule) {
//	r.nextRule=mr
//}

func (r *RuleTooMuchSteps) SetNextRule (tr TurnRule) {
	r.nextRule = tr
}

func (r *RuleAttemptToGetFewSticksFromHead) SetNextRule (tr TurnRule) {
	r.nextRule = tr
}

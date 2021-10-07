package game

type RuleMatchOrder struct {
	nextRule TurnRule
}

func (r *RuleMatchOrder) SetNextRule(gr TurnRule) {
	r.nextRule = gr
}

type RuleCorrectGamePhase struct {
	nextRule TurnRule
}

func (r *RuleCorrectGamePhase) SetNextRule(gr TurnRule) {
	r.nextRule = gr
}

type RuleMatchTurnNumber struct {
	nextRule TurnRule
}

func (r *RuleMatchTurnNumber) SetNextRule(gr TurnRule) {
	r.nextRule = gr
}

type RuleMoveMatchStickColor struct {
	nextRule MovingRule
}

func (r *RuleMoveMatchStickColor) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleMoveDirection struct {
	nextRule MovingRule
}

func (r *RuleMoveDirection) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleMoveImpossibleAmountSteps struct {
	nextRule MovingRule
}
func (r *RuleMoveImpossibleAmountSteps) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleMoveToOccupiedHole struct {
	nextRule MovingRule
}

func (r *RuleMoveToOccupiedHole) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleMoveFromEmptyHole struct {
	nextRule MovingRule
}

func (r *RuleMoveFromEmptyHole) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleForbiddenMoveKindLongBackgammon struct {
	nextRule MovingRule
}

type RuleMoveFormat struct {
	nextRule MovingRule
}

func (r *RuleMoveFormat) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleRemovingNotFromHome struct {
	nextRule MovingRule
}

func (r *RuleRemovingNotFromHome) SetNextRule(mr MovingRule) {
	r.nextRule = mr
}

type RuleTooMuchSteps struct {
	nextRule TurnRule
}

func (r *RuleTooMuchSteps) SetNextRule (tr TurnRule) {
	r.nextRule = tr
}

type RuleAttemptToGetFewSticksFromHead struct{
	nextRule TurnRule
}


func (r *RuleAttemptToGetFewSticksFromHead) SetNextRule (tr TurnRule) {
	r.nextRule = tr
}


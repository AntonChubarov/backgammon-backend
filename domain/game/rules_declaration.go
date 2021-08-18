package game

type RuleMoveMatchStickColor struct {
	nextRule MovingRule
}

type RuleMoveDirection struct {
	nextRule MovingRule
}

type RuleMoveImpossibleAmountSteps struct {
	nextRule MovingRule
}

type RuleMoveToOccupiedHole struct {
	nextRule MovingRule
}
package game

type RuleMatchOrder struct {
	nextRule GameRule
}

type RuleCorrectGamePhase struct {
	nextRule GameRule
}

type RuleMatchTurnNumber struct {
	nextRule GameRule
}

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

type RuleMoveFromEmptyHole struct {
	nextRule MovingRule
}

type RuleForbiddenMoveKindLongBackgammon struct {
	nextRule MovingRule
}

type RuleMoveFormat struct {
	nextRule MovingRule
}

type RuleRemovingNotFromHome struct {
	nextRule MovingRule
}

type RuleTooMuchSteps struct {
	nextRule TurnRule
}

type RuleAttemptToGetFewSticksFromHead struct{
	nextRule TurnRule
}
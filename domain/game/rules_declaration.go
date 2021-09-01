package game

type RuleMatchOrder struct {
	nextRule TurnRule
}

type RuleCorrectGamePhase struct {
	nextRule TurnRule
}

type RuleMatchTurnNumber struct {
	nextRule TurnRule
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
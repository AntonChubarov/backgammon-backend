package longbackgammon

//Rule001
func (r *RuleMatchOrder) ValidateRule(g *Game, t *Turn) error {
	if g.CurrentTurn != t.Color {
		return ErrorOutOfTurn
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, t)
	}
	return nil
}

//Rule002
func (r *RuleCorrectGamePhase) ValidateRule(g *Game, t *Turn) error {
	if g.State != InProcess {
		return ErrorOutOfGame
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, t)
	}
	return nil
}

func (r *RuleMatchTurnNumber) ValidateRule(g *Game, t *Turn) error {
	if g.AwaitingTurnNumber != t.TurnNumber {
		return ErrorInvalidTurnNumber
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, t)
	}
	return nil
}

//Rule003
func (r *RuleMoveMatchStickColor) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	if g.Board.Holes[m.From].Color != c && g.Board.Holes[m.From].StickCount != 0 {
		return ErrorOpponentsStickMoveAttempt
	}

	// Use this code everywhere, or make an iteration over the rules slice
	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

//Rule004
func (r *RuleMoveDirection) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	if c == White {
		if InvertNumeration(m.From) >= InvertNumeration(m.To) {
			return ErrorIncorrectMoveDirection
		}
	}
	if c == Black {
		if m.From >= m.To {
			return ErrorIncorrectMoveDirection
		}
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

//Rule005
func (r *RuleMoveImpossibleAmountSteps) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	//TODO THis is draft function! It ignores consumedDice
	if consumedDice != nil {
		if len(consumedDice) > 0 {
			panic("consumed Dice are not supported yet")
		}
	}

	distance := MoveDistance(c, m.From, m.To)
	if g.DiceState.Dice1 == g.DiceState.Dice2 {
		if distance == g.DiceState.Dice1 {
			return nil
		}
		if distance == g.DiceState.Dice1*2 {
			return nil
		}
		if distance == g.DiceState.Dice1*3 {
			return nil
		}
		if distance == g.DiceState.Dice1*4 {
			return nil
		}
	}
	if distance == g.DiceState.Dice1 {
		return nil
	}
	if distance == g.DiceState.Dice2 {
		return nil
	}
	if distance == g.DiceState.Dice1+g.DiceState.Dice2 {
		return nil
	}

	if r.nextRule == nil {
		return ErrorIncorrectNumberOfStepsInMove
	}
	return r.nextRule.ValidateRule(g, c, m, consumedDice)

}

//Rule006
func (r *RuleMoveToOccupiedHole) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	if g.Board.Holes[m.To].Color == -c {
		return ErrorMoveToOccupiedHole
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

//Rule008
func (r *RuleMoveFromEmptyHole) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	if g.Board.Holes[m.From].StickCount == 0 {
		return ErrorMoveFromEmptyHole
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

func (r *RuleForbiddenMoveKindLongBackgammon) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	if m.MoveKind == Movement {
		return nil
	}
	if m.MoveKind == Removing {
		return nil
	}
	if m.MoveKind == Surrender {
		return nil
	}

	if r.nextRule == nil {
		return ErrorImpossibleMoveKind
	}
	return r.nextRule.ValidateRule(g, c, m, consumedDice)

}

func (r *RuleMoveFormat) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	moveTypeCheckFail := m.MoveKind != Movement &&
		m.MoveKind != Removing &&
		m.MoveKind != Surrender

	if moveTypeCheckFail {
		return ErrorIncorrectMoveFormat
	}

	if m.MoveKind == Movement {
		if m.From < 1 || m.From > 24 || m.To < 1 || m.To > 24 {
			return ErrorIncorrectMoveFormat
		}
	}

	if m.MoveKind == Removing {
		if m.From < 1 || m.From > 24 || m.To != 0 {
			return ErrorIncorrectMoveFormat
		}
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

func (r *RuleRemovingNotFromHome) ValidateRule(g *Game, c Color, m *Move, consumedDice []int) error {
	if c == White {
		m.From = InvertNumeration(m.From)
	}
	if m.From < 19 {
		return ErrorRemovingFromInvalidHole
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

func (r *RuleTooMuchSteps) ValidateRule(g *Game, t *Turn) error {
	expectedStepsNumber := 2
	if g.Dice1 == g.Dice2 {
		expectedStepsNumber = 4
	}
	if len(t.Moves) > expectedStepsNumber {
		return ErrorTooMuchStepsInTurn
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, t)
	}
	return nil
}

func (r *RuleAttemptToGetFewSticksFromHead) ValidateRule(g *Game, t *Turn) error {
	headCount := 0
	if t.Color == Black {
		for i := range t.Moves {
			if t.Moves[i].From == 1 {
				headCount++
			}
		}
	}
	if t.Color == White {
		for i := range t.Moves {
			if t.Moves[i].From == 13 {
				headCount++
			}
		}
	}
	if headCount > 1 {
		return ErrorMoveFromHeadLimit1
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, t)
	}
	return nil
}

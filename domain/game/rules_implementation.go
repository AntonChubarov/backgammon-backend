package game

import (
	"backgammon/domain/board"
	"fmt"
	"log"
)

//Rule003
func (r *RuleMoveMatchStickColor) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	if g.Board.Holes[m.From].StickColor != c && g.Board.Holes[m.From].StickCount != 0 {
		return ErrorOpponentsStickMoveAttempt
	}

	if r.nextRule != nil {
		return r.nextRule.ValidateRule(g, c, m, consumedDice)
	}
	return nil
}

//Rule004
func (r *RuleMoveDirection) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {

	switch c {
	case board.White:
		if board.InvertNumeration(m.From)>=board.InvertNumeration(m.To) {return ErrorIncorrectMoveDirection}
	case board.Black:
		if m.From>=m.To {return ErrorIncorrectMoveDirection}
	default:
		err:=fmt.Errorf("unexpected stick color %d in RuleMoveDirection ValidateRule", c)
		log.Println(err)
		return err
	}

	if r.nextRule!=nil {return r.nextRule.ValidateRule(g, c, m, consumedDice)}
	return nil
}

//Rule005
func (r *RuleMoveImpossibleAmountSteps) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	//TODO THis is draft function! It ignores consumedDice
	if consumedDice!=nil {
		if len(consumedDice)>0 {panic("consumed Dice are not supported yet")}
	}

	distance:=MoveDistance(c, m.From, m.To)
	if g.DiceState.Dice1 == g.DiceState.Dice2 {
		if distance==g.DiceState.Dice1 {return nil}
		if distance==g.DiceState.Dice1*2 {return nil}
		if distance==g.DiceState.Dice1*3 {return nil}
		if distance==g.DiceState.Dice1*4 {return nil}
	}
	if distance==g.DiceState.Dice1 {return nil}
	if distance==g.DiceState.Dice2 {return nil}
	if distance==g.DiceState.Dice1+g.DiceState.Dice2 {return nil}

	if r.nextRule==nil {return ErrorIncorrectNumberOfStepsInMove}
	return r.nextRule.ValidateRule(g, c, m, consumedDice)

}

//Rule006
func (r *RuleMoveToOccupiedHole) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	if g.Board.Holes[m.To].StickColor==-c { return ErrorMoveToOccupiedHole}

	if r.nextRule!=nil {return r.nextRule.ValidateRule(g, c, m, consumedDice)}
	return nil
}
//Rule008
func (r *RuleMoveFromEmptyHole) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	if g.Board.Holes[m.From].StickCount==0 { return ErrorMoveFromEmptyHole}

	if r.nextRule!=nil {return r.nextRule.ValidateRule(g, c, m, consumedDice)}
	return nil
}

func (r *RuleForbiddenMoveKindLongBackgammon) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	if m.MoveKind==board.Movement {return nil}
	if m.MoveKind==board.Removing {return nil}
	if m.MoveKind==board.Surrender {return nil}

	if r.nextRule==nil {return ErrorImpossibleMoveKind}
	return r.nextRule.ValidateRule(g, c, m, consumedDice)

}


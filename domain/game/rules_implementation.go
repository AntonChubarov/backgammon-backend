package game

import (
	"backgammon/domain/board"
	"fmt"
	"log"
)

//Rule003
func (r *RuleMoveMatchStickColor) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	if g.Board.Holes[m.From].StickColor !=c { return ErrorOpponentsStickMoveAttempt}

	if r.nextRule!=nil {return r.nextRule.ValidateRule(g, c, m, consumedDice)}
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

	//TODO - implement method!

	if r.nextRule!=nil {return r.nextRule.ValidateRule(g, c, m, consumedDice)}
	return nil
}


//Rule006
func (r *RuleMoveToOccupiedHole) ValidateRule(g *Game, c board.StickColor, m *board.Move, consumedDice []int) error {
	if g.Board.Holes[m.To].StickColor==-c { return ErrorMoveToOccupiedHole}

	if r.nextRule!=nil {return r.nextRule.ValidateRule(g, c, m, consumedDice)}
	return nil
}


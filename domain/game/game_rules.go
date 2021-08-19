package game

import "backgammon/domain/board"

type RulesKeeper interface {
	InitGame(game *Game)
	PerformTurn(game *Game, stickColor *board.StickColor, turn *board.Turn) error
	GetDiceInterpretation() func(d *board.DiceState) []int
}

func DiceInterpretationLongBackgammon(d *board.DiceState) []int {
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

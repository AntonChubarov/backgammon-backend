package game

import (
	"backgammon/domain/board"
)

type State int

const (
	NotStarted State = iota
	InProcess
	Finished
	StoppedError
)

type Winner int

const (
	None Winner = iota
	White
	Black
	Draw
)

type Game struct {
	State
	Winner
	board.Board
	board.DiceState
	RulesKeeper
	CurrentTurn board.StickColor
	AwaitingTurnNumber int
}

type RulesKeeper interface {
	InitGame(game *Game)
	PerformTurn(game *Game, stickColor *board.StickColor, turn *board.Turn) error
	GetDiceInterpretation() func(d *board.DiceState) []int
	ElectFirstMove() board.StickColor
}





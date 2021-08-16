package game

import (
	"backgammon/domain/board"
)

type State int

const (
	NotStarted State = iota
	InProcess
	Finished
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




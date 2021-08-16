package game

import "backgammon/domain/board"

type RulesKeeper interface {
	InitGame(game *Game)
	PerformTurn(game *Game, stickColor *board.StickColor, turn *board.Turn) error

}

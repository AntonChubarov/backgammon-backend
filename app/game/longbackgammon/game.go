package longbackgammon

type State int

const (
	NotStarted State = iota
	InProcess
	Finished
	StoppedError
)

//type Winner int
//
//const (
//	None Winner = iota
//	White
//	Black
//	Draw
//)

type Game struct {
	State
	//	Winner
	Board
	DiceState
	RulesKeeper
	CurrentTurn        Color
	AwaitingTurnNumber int
}

type RulesKeeper interface {
	InitGame(game *Game)
	PerformTurn(game *Game, turn *Turn) error
	GetDiceInterpretation() func(d *DiceState) []int
	ElectFirstMove() Color
}

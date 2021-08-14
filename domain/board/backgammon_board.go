package board

type StickColor int

const (
	Black StickColor = iota - 1
	_
	White
)
type MoveKind int

const (
	_ MoveKind = iota
	Movement
	Removing
)



type Hole struct {
	StickColor     // -1 is black color, 1 is white color
	StickCount int // should be an integer between 0 and 15
}

type Turn struct {
	StickColor //-1 is black color, 1 is white color
	Moves []Move   // can contain from 0 to 4 moves
}

type Move struct {
	MoveKind
	From int // should be an integer between 1 and 24
	Steps int // which way we should determine stick remove?
}

type Board struct {
	CurrentDiceState DiceState
	CurrentTurnColor StickColor
	Holes [24]Hole
}

func (b *Board) Clear() {
	b.CurrentTurnColor=White
	b.CurrentDiceState=DiceState{Dice1: 0, Dice2: 0}
	for i:= range b.Holes {
		b.Holes[i].StickCount =0
	}
}


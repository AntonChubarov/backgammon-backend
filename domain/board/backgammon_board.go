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
	Surrender = 999
)



type Hole struct {
	StickColor     // -1 is black color, 1 is white color
	StickCount int // should be an integer between 0 and 15
}

type Turn struct {
	StickColor //-1 is black color, 1 is white color
	Moves []Move   // can contain from 0 to 4 moves
	TurnNumber int
}

type Move struct {
	MoveKind
	From int
	To int
}

type Board struct {
	Holes [25]Hole //Hole[0] - empty placeholder
}

func (b *Board) Clear() {
	for i:= range b.Holes {
		b.Holes[i].StickCount =0
	}
}

func InvertNumeration(n int) int {
	if n<=12 {return n +12}
	return n-12
}



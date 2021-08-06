package board

import "fmt"

var ImpossibleMove = fmt.Errorf("this move is impossible")
var IncompleteMove = fmt.Errorf("this move is incomplete")

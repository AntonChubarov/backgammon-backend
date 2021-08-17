package game

import "fmt"

// (+) 1. Ходит не в свою очередь
// (+) 2. Ход вне игры (до начала, или после конца, или ???)
// -----ошибки движений
// (+)3. Попытка походить не своей фишкой
// (+)4. Попытка походить не в том направлении (или с пересечением границы поля)
// (+)5. Попытка походить на невозможное количество шагов
// (+)6. Попытка походить в занятую ячейку
// (+)7. Попытка снятия больше 1 фишки с головы
// (+)8. Попытка хода походить из пустой ячейки
// (+)9. Попытка выполнить движение, запрещенное в этом типе игры
// .. по одному "запрещателю" на каждый тип движения
// (+)10. Попытка выполнить движение, запрещенное в этой игровой фазе
// .. по одному "запрещателю" на каждый тип движения
// (+)11. Неправильный формат движения
// .. по одному "запрещателю" на каждый тип движения

// ---- ошибки хода
// 1. Забор, запирающий все фишки соперника
// 2. Неполный ход (Ход на немаксимальное из возможных количество шагов)
// 3. Номер хода не соответствует ожидаемому

var ErrorOutOfTurn = fmt.Errorf("attempt to make move out of turn")
var ErrorUotOfGame = fmt.Errorf("attemt to make move out of game")

var ErrorOpponentsStickMoveAttempt = fmt.Errorf("attempt to make move by opponent's stick")
var ErrorIncorrectMoveDirection = fmt.Errorf("attemt to make move in wrong direction")
var ErrorIncorrectNumberOfSteps = fmt.Errorf("attempt to make move on wrong number of steps")
var ErrorMoveToOccupiedHole = fmt.Errorf("attempt to make move in a hole, ocupied by opponent")
var ErrorMoveFromHeadLimit1 = fmt.Errorf("attempt to use more than 1 stick from head due turn")
var ErrorMoveFromEmptyHole = fmt.Errorf("attempt to make move from empty hole")

var ErrorImpossibleMoveKind = fmt.Errorf("attempt to make move, which type is disallowed in curent game type")

var ErrorImpossibleMoveKindInGamePhase = fmt.Errorf("attempt to make move, which type is disallowed in current game phase")

var ErrorIncorrectMoveFormat = fmt.Errorf("incorrect move data received")

var ErrorBlockingFence  = fmt.Errorf("attempt to build blocking fence at the end of turn")
var ErrorIncompleteTurn = fmt.Errorf("attempt to make incomplete turn")
var ErrorIncorrectTurnSerialNumber = fmt.Errorf("turn with incorrect serial number received")







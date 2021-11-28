package game

import (
	"backgammon/app/game/longbackgammon"
	"backgammon/domain"
	"backgammon/utils"
	"github.com/gorilla/websocket"
)

// TODO Bring an order

//У кого должен быть метод "игрок пытается сделать ход?"
//Кому и через какой метод, мы должны сообщить в конечном счете игроку, что ход получилось сделать?
//Кому и через какой метод, мы должны сообщить, что изменилось состояние игры,
//	в частности, расстановка фишек, состояние кубиков и т.д.

//Кому и через какой метод, мы должны послать сообщение, что приглашаем игрока сделать ход?

type PlayerBase struct {
	CurrentRoom *Room
	token       string
	//TODO change websocket to abstract (interface) communicator
	ws *websocket.Conn
}

type PlayerInterface interface { //TODO change to normal name
	SendGameState(g *longbackgammon.Game)
	InviteMakeMove()
}

type Spectator interface {
	SendGameState(g *longbackgammon.Game)
}
type Room struct {
	Players []PlayerInterface
	longbackgammon.Game
	Spectators []Spectator
	UUID       domain.UUID
	Name       string
	LobbyInterface
}

func NewRoom(game longbackgammon.Game, name string) *Room {
	return &Room{Game: game,
		UUID: utils.GenerateUUID(),
		Name: name}
}

type RoomInterface interface {
	MakeTurn(t *longbackgammon.Turn) error
	//Must return Player, sitting in certain Room
	AddPlayer(player PlayerInterface, c longbackgammon.Color) error
	//When room is full of players, it starts game, invoking it method ("StartGame"(??))
	AddSpectator(spectator Spectator) error
}

type RoomStorage interface {
	Add(r *Room)
	GetRoomByUUID(uuid string) *Room
	RemoveByUUID(uuid string) error
	GetRooms(pageSize int, page int) []*Room
}

type LobbyBase struct {
	RoomStorage
}

type LobbyInterface interface {
	CreateHumanPlayer(token string, ws *websocket.Conn) PlayerInterface //Creates player, based on websocketConnection
	//and active usersession
	CreateRoom(game longbackgammon.Game, name string, creatorPlayer PlayerInterface) *Room //Creates Room, and adds it to storage

	AddPlayerToRoomByUUID(roomUUID string, player PlayerInterface) error
}

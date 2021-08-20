package auth

import "github.com/gorilla/websocket"

type SessionStorage interface {
	AddNewUser(data *UserSessionData)
	UpdateTokenExpiryTime(token string)
	DeleteUserByToken(token string)
	GetTokenByUUID(uuid string) (token string, wasFound bool)
	IsTokenValid(token string) bool
	//TODO REF remove to other entity
	SetWebSocketToUserByToken(token string, webSocket *websocket.Conn)
	//TODO
	//GetUserUUIDByToken(token string) (UUID string, wasFound bool)
}

type StringHasher interface {
	HashString(s string) (string, error)
}

type TokenGenerator interface {
	GenerateToken () string
}


package auth

import (
	"github.com/gorilla/websocket"
	"time"
)

type UserAuthData struct {
	UUID string
	Login string
	Password string
	Token string
}

type UserSessionData struct {
	Token string
	ExpiryTime time.Time
	UserUUID string
	WebSocket websocket.Conn
}

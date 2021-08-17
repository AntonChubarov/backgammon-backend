package auth

import (
	"github.com/gorilla/websocket"
	"time"
)

type UserSessionData struct {
	Token string
	ExpiryTime time.Time
	UserUUID string
	WebSocket websocket.Conn
}




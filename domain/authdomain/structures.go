package authdomain

import "time"

type UUID string
type Token string
type UserName string
type Password string
type RoomID string
type ExpiryTime time.Time


type UserData struct {
	UUID
	UserName
	Password
}



type UserSessionData struct {
	UUID
	Token
	RoomID
	ExpiryTime
}

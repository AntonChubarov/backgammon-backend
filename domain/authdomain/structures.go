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

func (u *UserData) Update(sample *UserData) {
	u.UserName=sample.UserName
	u.Password=sample.Password
}

type UserSessionData struct {
	UUID
	Token
	RoomID
	ExpiryTime
}

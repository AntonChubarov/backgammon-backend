package domain

import "time"

// TODO move dto's to specific packages

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Lobby DTO structures

type RoomsInfoDTO struct {
	RoomsInfo []RoomInfoDTO `json:"rooms_info"`
	Message   string        `json:"message"`
}

type RoomInfoDTO struct {
	RoomName        string `json:"room_name"`
	WhitePlayerName string `json:"white_player_name"`
	BlackPlayerName string `json:"black_player_name"`
}

type CreateRoomRequestDTO struct {
	Token    string `json:"token"`
	RoomName string `json:"room_name"`
	Color    int    `json:"color"`
}

type ConnectToRoomRequestDTO struct {
	Token    string `json:"token"`
	RoomName string `json:"room_name"`
}

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

type SessionData struct {
	UUID
	Token
	RoomID
	ExpiryTime
}

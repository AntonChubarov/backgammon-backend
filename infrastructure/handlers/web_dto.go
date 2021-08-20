package handlers

// Auth DTO structures

type UserAuthRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistrationResponseDTO struct {
	Message string `json:"message"`
}

type UserAuthorizationResponseDTO struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

// Lobby DTO structures

type RoomsInfoDTO struct {
	RoomsInfo []RoomInfoDTO `json:"rooms_info"`
	Message string `json:"message"`
}

type RoomInfoDTO struct {
	RoomName string `json:"room_name"`
	WhitePlayerName string `json:"white_player_name"`
	BlackPlayerName string `json:"black_player_name"`
}

type CreateRoomRequestDTO struct {
	Token string `json:"token"`
	RoomName string `json:"room_name"`
	Color int `json:"color"`
}

type CreateRoomResponseDTO struct {
	Message string `json:"message"`
}

type ConnectToRoomRequestDTO struct {
	Token string `json:"token"`
	RoomName string `json:"room_name"`
}

type ConnectToRoomResponseDTO struct {
	Message string `json:"message"`
}

type DisconnectRequestDTO struct {
	Token string `json:"token"`
}

type DisconnectResponseDTO struct {
	Message string `json:"message"`
}
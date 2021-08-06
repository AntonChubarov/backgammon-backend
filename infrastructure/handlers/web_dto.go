package handlers

type UserRegistrationRequestDTO struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type UserRegistrationResponseDTO struct {
	Message string `json:"message"`
}

type UserAuthorizationResponseDTO struct {
	Message string `json:"message"`
	Token string `json:"token"`
}
package handlers

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
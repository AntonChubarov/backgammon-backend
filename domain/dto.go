package domain

type UserRegistrationRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type UserRegistrationResponse struct {
	Message string `json:"message"`
}

type UserAuthorizationResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

type UserDBDTO struct {
	Login string `db:"login"`
	PasswordHash string `db:"password"`
}
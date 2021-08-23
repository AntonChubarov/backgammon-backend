package auth

import (
	"backgammon/domain/authdomain"
)

func UserDataToUserDBDTO(user authdomain.UserAuthData) UserDBDTO {
	return UserDBDTO{
		UUID:         user.UUID,
		Username:     user.Username,
		PasswordHash: user.Password,
	}
}

func UserDBDTOToUserData(user UserDBDTO) authdomain.UserAuthData {
	return authdomain.UserAuthData{
		UUID:     user.UUID,
		Username: user.Username,
		Password: user.PasswordHash,
	}
}
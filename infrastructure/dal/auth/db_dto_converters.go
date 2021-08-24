package auth

import (
	"backgammon/domain/authdomain"
)

func UserDataToUserDBDTO(user authdomain.UserData) UserDBDTO {
	return UserDBDTO{
		UUID:         string(user.UUID),
		Username:     string(user.UserName),
		PasswordHash: string(user.Password),
	}
}

func UserDBDTOToUserData(user UserDBDTO) authdomain.UserData {
	return authdomain.UserData{
		UUID:     authdomain.UUID(user.UUID),
		UserName: authdomain.UserName(user.Username),
		Password: authdomain.Password(user.PasswordHash),
	}
}
package dal

import (
	"backgammon/domain/auth"
)

func UserDataToUserDBDTO(user auth.UserAuthData) UserDBDTO {
	return UserDBDTO{
		UUID:         user.UUID,
		Username:     user.Username,
		PasswordHash: user.Password,
	}
}

func UserDBDTOToUserData(user UserDBDTO) auth.UserAuthData {
	return auth.UserAuthData{
		UUID:     user.UUID,
		Username: user.Username,
		Password: user.PasswordHash,
	}
}
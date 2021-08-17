package dal

import (
	"backgammon/domain/auth"
)

func UserDataToUserDBDTO(user auth.UserAuthData) UserDBDTO {
	return UserDBDTO{
		UUID:         user.UUID,
		Login:        user.Login,
		PasswordHash: user.Password,
	}
}

func UserDBDTOToUserData(user UserDBDTO) auth.UserAuthData {
	return auth.UserAuthData{
		UUID: user.UUID,
		Login: user.Login,
		Password: user.PasswordHash,
	}
}
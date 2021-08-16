package dal

import (
	"backgammon/domain"
)

func UserDataToUserDBDTO(user domain.UserAuthData) UserDBDTO {
	return UserDBDTO{
		UUID:         user.UUID,
		Login:        user.Login,
		PasswordHash: user.Password,
	}
}

func UserDBDTOToUserData(user UserDBDTO) domain.UserAuthData {
	return domain.UserAuthData{
		UUID: user.UUID,
		Login: user.Login,
		Password: user.PasswordHash,
	}
}
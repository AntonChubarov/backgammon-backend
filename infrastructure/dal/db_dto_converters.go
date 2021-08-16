package dal

import (
	"backgammon/domain"
)

func UserDataToUserDBDTO(user domain.UserAuthData) UserDBDTO {
	return UserDBDTO{
		Id: user.UUID,
		Login:        user.Login,
		PasswordHash: user.Password,
	}
}

func UserDBDTOToUserData(user UserDBDTO) domain.UserAuthData {
	return domain.UserAuthData{
		UUID: user.Id,
		Login: user.Login,
		Password: user.PasswordHash,
	}
}
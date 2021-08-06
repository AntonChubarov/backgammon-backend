package dal

import (
	"backgammon/domain"
)

func UserDataToUserDBDTO(user domain.UserData) UserDBDTO {
	return UserDBDTO{
		Id: user.UUID,
		Login:        user.Login,
		PasswordHash: user.Password,
	}
}

func UserDBDTOToUserData(user UserDBDTO) domain.UserData {
	return domain.UserData{
		UUID: user.Id,
		Login: user.Login,
		Password: user.PasswordHash,
	}
}
package dal

import (
	"backgammon/domain"
)

func UserToDatabase(user domain.UserData) UserDBDTO {
	return UserDBDTO{
		Login:        user.Login,
		PasswordHash: user.Password,
	}
}

func UserFromDatabase(user UserDBDTO) domain.UserData {
	return domain.UserData{
		Login: user.Login,
		Password: user.PasswordHash,
	}
}
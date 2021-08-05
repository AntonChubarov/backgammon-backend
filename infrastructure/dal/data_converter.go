package dal

import (
	"backgammon/domain"
)

func UserToDatabase(data domain.UserData) domain.UserDBDTO {
	return domain.UserDBDTO{
		Login: data.Login,
		PasswordHash: data.Password,
	}
}
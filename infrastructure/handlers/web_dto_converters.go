package handlers

import (
	"backgammon/domain/auth"
)

func ConvertUserRegDataToUser(user UserAuthRequestDTO) auth.UserAuthData {
	return auth.UserAuthData{
		Login: user.Login,
		Password: user.Password,
		Token: "",
	}
}
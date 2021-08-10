package handlers

import (
	"backgammon/domain"
)

func ConvertUserRegDataToUser(user UserAuthRequestDTO) domain.UserData {
	return domain.UserData{
		Login: user.Login,
		Password: user.Password,
		Token: "",
	}
}
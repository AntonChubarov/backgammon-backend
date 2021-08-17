package handlers

import (
	"backgammon/domain"
)

func ConvertUserRegDataToUser(user UserAuthRequestDTO) domain.UserAuthData {
	return domain.UserAuthData{
		Login: user.Login,
		Password: user.Password,
		Token: "",
	}
}
package handlers

import (
	"backgammon/domain"
)

func ConvertUserRegDataToUser(user UserRegistrationRequestDTO) domain.UserData {
	return domain.UserData{
		Login: user.Login,
		Password: user.Password,
		Token: "",
	}
}
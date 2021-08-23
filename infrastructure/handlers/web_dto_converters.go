package handlers

import (
	"backgammon/domain/authdomain"
)

func ConvertUserRegDataToUser(user UserAuthRequestDTO) authdomain.UserAuthData {
	return authdomain.UserAuthData{
		Username: user.Username,
		Password: user.Password,
		Token:    "",
	}
}
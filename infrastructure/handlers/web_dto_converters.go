package handlers

import (
	"backgammon/domain/auth"
)

func ConvertUserRegDataToUser(user UserAuthRequestDTO) auth.UserAuthData {
	return auth.UserAuthData{
		Username: user.Username,
		Password: user.Password,
		Token:    "",
	}
}
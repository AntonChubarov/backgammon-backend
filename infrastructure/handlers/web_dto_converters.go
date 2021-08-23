package handlers

import (
	"backgammon/domain/authdomain"
)

func ConvertUserRegDataToUser(user UserAuthRequestDTO) authdomain.UserData {
	return authdomain.UserData{
		UUID: "",
		UserName: authdomain.UserName(user.Username),
		Password: authdomain.Password(user.Password),
	}
}
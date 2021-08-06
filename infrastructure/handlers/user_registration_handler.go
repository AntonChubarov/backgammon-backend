package handlers

import (
	"backgammon/app"
	"backgammon/domain"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserRegistrator struct {
	service *app.UserRegistrationService
}

func NewUserRegistrator(service *app.UserRegistrationService) *UserRegistrator {
	return &UserRegistrator{service: service}
}

func (ur *UserRegistrator) Handle(c echo.Context) error {
	var err error
	var request UserRegistrationRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.UserRegistrator", err)
		return echo.ErrBadRequest
	}

	user := ConvertUserRegDataToUser(request)
	err = ur.service.RegisterNewUser(user)

	if err == domain.UserExistError {
		errStr := err.Error()
		return c.JSON(http.StatusConflict, UserRegistrationResponseDTO{Message: errStr})
	}

	if err != nil {
		log.Println("In handlers.UserRegistrator", err)
		return c.JSON(http.StatusInternalServerError, UserRegistrationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserRegistrationResponseDTO{Message: "Successfully registered"})
}
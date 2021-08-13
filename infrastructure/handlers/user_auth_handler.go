package handlers

import (
	"backgammon/app"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserAuthHandler struct {
	service *app.UserAuthService
}

func NewUserAuthHandler(service *app.UserAuthService) *UserAuthHandler {
	return &UserAuthHandler{service: service}
}

func (uah *UserAuthHandler) Register(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.UserAuthHandler.Register", err)
		return echo.ErrBadRequest
	}

	user := ConvertUserRegDataToUser(request)
	err = uah.service.RegisterNewUser(user)

	if err == app.UserExistError {
		errStr := err.Error()
		return c.JSON(http.StatusConflict, UserRegistrationResponseDTO{Message: errStr})
	}

	if err != nil {
		log.Println("In handlers.UserAuthHandler.Register", err)
		return c.JSON(http.StatusInternalServerError, UserRegistrationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserRegistrationResponseDTO{Message: "Successfully registered"})
}

func (uah *UserAuthHandler) Login(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.UserAuthHandler.Login", err)
		return echo.ErrBadRequest
	}

	user := ConvertUserRegDataToUser(request)
	// Authorize User
	var token string
	token, err = uah.service.AuthorizeUser(user)

	if err == app.InvalidLogin || err == app.InvalidPassword {
		errStr := err.Error()
		return c.JSON(http.StatusUnauthorized, UserAuthorizationResponseDTO{Message: errStr})
	}

	if err != nil {
		log.Println("In handlers.UserAuthHandler.Login", err)
		return c.JSON(http.StatusInternalServerError, UserAuthorizationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserAuthorizationResponseDTO{Message: "Authorized", Token: token})
}
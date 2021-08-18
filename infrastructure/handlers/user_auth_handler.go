package handlers

import (
	"backgammon/app/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserAuthHandler struct {
	service *auth.UserAuthService
}

func NewUserAuthHandler(service *auth.UserAuthService) *UserAuthHandler {
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

	if err == auth.ErrorUserExists {
		errStr := err.Error()
		return c.JSON(http.StatusConflict, UserRegistrationResponseDTO{Message: errStr})
	}

	if err != nil {
		log.Println("In handlers.UserAuthHandler.Register", err)
		return c.JSON(http.StatusInternalServerError, UserRegistrationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserRegistrationResponseDTO{Message: "Successfully registered"})
}

func (uah *UserAuthHandler) Authorize(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.UserAuthHandler.Authorize", err)
		return echo.ErrBadRequest
	}

	user := ConvertUserRegDataToUser(request)
	// Authorize User
	var token string
	token, err = uah.service.AuthorizeUser(user)

	if err == auth.ErrorUserNotRegistered || err == auth.ErrorInvalidPassword {
		errStr := err.Error()
		return c.JSON(http.StatusUnauthorized, UserAuthorizationResponseDTO{Message: errStr})
	}

	if err != nil {
		log.Println("In handlers.UserAuthHandler.Authorize", err)
		return c.JSON(http.StatusInternalServerError, UserAuthorizationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserAuthorizationResponseDTO{Message: "Authorized", Token: token})
}
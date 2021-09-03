package handlers

import (
	"backgammon/app/auth"
	"backgammon/domain/authdomain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAuthHandler struct {
	userAuthService *auth.UserAuthService
}

func NewUserAuthHandler(service *auth.UserAuthService) *UserAuthHandler {
	return &UserAuthHandler{userAuthService: service}
}

func (uah *UserAuthHandler) Register(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Register", err)
		return echo.ErrBadRequest
	}

	//log.Println("register request from:", request.Username)

	user := ConvertUserRegDataToUser(request)
	err = uah.userAuthService.RegisterNewUser(user)

	if err == auth.ErrorUserExists {
		errStr := err.Error()
		return c.JSON(http.StatusConflict, UserRegistrationResponseDTO{Message: errStr})
	}

	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Register", err)
		return c.JSON(http.StatusInternalServerError, UserRegistrationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserRegistrationResponseDTO{Message: "Successfully registered"})
}

func (uah *UserAuthHandler) Authorize(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Authorize", err)
		return echo.ErrBadRequest
	}

	//log.Println("authorize request from:", request.Username)

	user := ConvertUserRegDataToUser(request)

	token, err := uah.userAuthService.AuthorizeUser(user)

	if err == auth.ErrorUserNotRegistered || err == auth.ErrorInvalidPassword {
		errStr := err.Error()
		return c.JSON(http.StatusUnauthorized, UserAuthorizationResponseDTO{Message: errStr})
	}

	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Authorize", err)
		return c.JSON(http.StatusInternalServerError, UserAuthorizationResponseDTO{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserAuthorizationResponseDTO{Message: "Authorized", Token: string(token)})
}

func (uah *UserAuthHandler) ProlongToken(c echo.Context) error {

	panic("Implement me")
}

func (uah *UserAuthHandler) Logout(c echo.Context) error {

	panic("Implement me")
}

func ConvertUserRegDataToUser(user UserAuthRequestDTO) authdomain.UserData {
	return authdomain.UserData{
		UUID:     "",
		UserName: authdomain.UserName(user.Username),
		Password: authdomain.Password(user.Password),
	}
}

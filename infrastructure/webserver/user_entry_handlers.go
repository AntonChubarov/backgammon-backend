package webserver

import (
	"backgammon/app/service"
	"backgammon/domain"
	"github.com/jbrodriguez/mlog"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserAuthHandler struct {
	userAuthService *service.UserEntryService
}

func NewUserAuthHandler(service *service.UserEntryService) *UserAuthHandler {
	return &UserAuthHandler{userAuthService: service}
}

func (uah *UserAuthHandler) Register(c echo.Context) error {
	var err error
	var request domain.UserCredentials

	err = c.Bind(&request)
	if err != nil {
		mlog.Error(err)
		return echo.ErrBadRequest
	}

	mlog.Info("register request from: %s", request.Username)

	user := ConvertUserRegDataToUser(request)
	err = uah.userAuthService.RegisterNewUser(user)

	if err == service.ErrorUserExists {
		errStr := err.Error()
		return c.JSON(http.StatusConflict, errStr)
	}

	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Register", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Successfully registered")
}

func (uah *UserAuthHandler) Authorize(c echo.Context) error {
	var err error
	var request domain.UserCredentials

	err = c.Bind(&request)
	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Authorize", err)
		return echo.ErrBadRequest
	}

	//log.Println("authorize request from:", request.Username)

	user := ConvertUserRegDataToUser(request)

	token, err := uah.userAuthService.AuthorizeUser(user)

	if err == service.ErrorUserNotRegistered || err == service.ErrorInvalidPassword {
		errStr := err.Error()
		return c.JSON(http.StatusUnauthorized, errStr)
	}

	if err != nil {
		//log.Println("In handlers.UserAuthHandler.Authorize", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, string(token))
}

func (uah *UserAuthHandler) ProlongToken(c echo.Context) error {

	panic("Implement me")
}

func (uah *UserAuthHandler) Logout(c echo.Context) error {

	panic("Implement me")
}

func ConvertUserRegDataToUser(user domain.UserCredentials) domain.UserData {
	return domain.UserData{
		UUID:     "",
		UserName: domain.UserName(user.Username),
		Password: domain.Password(user.Password),
	}
}

package webserver

import (
	"backgammon/app/service"
	"backgammon/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LobbyHandler struct {
	userAuthService *service.UserEntryService
}

func NewLobbyHandler(userAuthService *service.UserEntryService) *LobbyHandler {
	return &LobbyHandler{userAuthService: userAuthService}
}

func (lh *LobbyHandler) GetRoomsInfo(c echo.Context) error {
	var err error

	token := domain.Token(c.QueryParam("token"))

	err = lh.userAuthService.CheckToken(token)

	if err != nil {
		//log.Println("In handlers.LobbyHandler.GetRoomsInfo", err)
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, domain.RoomsInfoDTO{
		RoomsInfo: []domain.RoomInfoDTO{},
		Message:   "No rooms yet",
	})
}

func (lh *LobbyHandler) CreateRoom(c echo.Context) error {
	var err error
	var request domain.UserCredentials

	err = c.Bind(&request)
	if err != nil {
		//log.Println("In handlers.LobbyHandler.CreateRoom", err)
		return echo.ErrBadRequest
	}

	panic("Implement me")
}

func (lh *LobbyHandler) ConnectToRoom(c echo.Context) error {
	var err error
	var request domain.UserCredentials

	err = c.Bind(&request)
	if err != nil {
		//log.Println("In handlers.LobbyHandler.ConnectToRoom", err)
		return echo.ErrBadRequest
	}

	panic("Implement me")
}

func (lh *LobbyHandler) Disconnect(c echo.Context) error {
	var err error
	var request domain.UserCredentials

	err = c.Bind(&request)
	if err != nil {
		//log.Println("In handlers.LobbyHandler.Disconnect", err)
		return echo.ErrBadRequest
	}

	panic("Implement me")
}

package handlers

import (
	"backgammon/app/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type LobbyHandler struct {
	userAuthService *auth.UserAuthService
}

func NewLobbyHandler(userAuthService *auth.UserAuthService) *LobbyHandler {
	return &LobbyHandler{userAuthService: userAuthService}
}

func (lh *LobbyHandler) GetRoomsInfo(c echo.Context) error {
	var err error

	token := c.QueryParam("token")

	err = lh.userAuthService.CheckToken(token)

	if err != nil {
		log.Println("In handlers.LobbyHandler.GetRoomsInfo", err)
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, RoomsInfoDTO{
		RoomsInfo: []RoomInfoDTO{},
		Message: "No rooms yet",
	})
}

func (lh *LobbyHandler) CreateRoom(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.LobbyHandler.CreateRoom", err)
		return echo.ErrBadRequest
	}


	panic("Implement me")
}

func (lh *LobbyHandler) ConnectToRoom(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.LobbyHandler.ConnectToRoom", err)
		return echo.ErrBadRequest
	}


	panic("Implement me")
}

func (lh *LobbyHandler) Disconnect(c echo.Context) error {
	var err error
	var request UserAuthRequestDTO

	err = c.Bind(&request)
	if err != nil {
		log.Println("In handlers.LobbyHandler.Disconnect", err)
		return echo.ErrBadRequest
	}


	panic("Implement me")
}
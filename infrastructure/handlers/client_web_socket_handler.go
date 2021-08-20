package handlers

import (
	"backgammon/app/auth"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebSocketHandler struct {
	webSocketManageService *auth.WebSocketManageService
}

func NewWebSocketHandler(webSocketManageService *auth.WebSocketManageService) *WebSocketHandler {
	return &WebSocketHandler{webSocketManageService: webSocketManageService}
}

func (wsh *WebSocketHandler) Handle(c echo.Context) (err error) {
	token := c.QueryParam("token")

	err = wsh.webSocketManageService.CheckToken(token)
	if err != nil {
		return err
	}

	upgrader := websocket.Upgrader{}
	var ws *websocket.Conn
	ws, err = upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	wsh.webSocketManageService.SetWebSocketByToken(token, ws)

	return
}
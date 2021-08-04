package websockethandler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebSocketHandler struct {
	
}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{}
}

func (wsh *WebSocketHandler) Handle(c echo.Context) error {
	upgrader := websocket.Upgrader{}

	var err error
	var ws *websocket.Conn

	ws, err = upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		err = ws.WriteJSON("message")
		if err != nil {
			c.Logger().Error(err)
		}

		var msg string
		err = ws.ReadJSON(&msg)
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
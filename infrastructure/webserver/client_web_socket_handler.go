package webserver

//import (
//	"backgammon/app/service"
//	"github.com/gorilla/websocket"
//	"github.com/labstack/echo/v4"
//	"log"
//)
//
//type WebSocketHandler struct {
//	userAuthService *service.UserAuthService
//	webSocketManageService *service.WebSocketManageService
//}
//
//func NewWebSocketHandler(userAuthService *service.UserAuthService, webSocketManageService *service.WebSocketManageService) *WebSocketHandler {
//	return &WebSocketHandler{userAuthService: userAuthService, webSocketManageService: webSocketManageService}
//}
//
//func (wsh *WebSocketHandler) Handle(c echo.Context) (err error) {
//	token := c.QueryParam("token")
//
//	err = wsh.userAuthService.CheckToken(token)
//	if err != nil {
//		return err
//	}
//
//	log.Println("websocket request with token:", token)
//
//	upgrader := websocket.Upgrader{}
//	var ws *websocket.Conn
//	ws, err = upgrader.Upgrade(c.Response(), c.Request(), nil)
//	if err != nil {
//		return err
//	}
//
//	wsh.webSocketManageService.SetWebSocketByToken(token, ws)
//
//	return
//}

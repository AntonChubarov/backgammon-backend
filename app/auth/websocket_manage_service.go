package auth

import (
	"backgammon/domain/authdomain"
)

type WebSocketManageService struct {
	mainSessionStorage authdomain.SessionStorage
}

func NewWebSocketManageService(mainSessionStorage authdomain.SessionStorage) *WebSocketManageService {
	return &WebSocketManageService{mainSessionStorage: mainSessionStorage}
}

//func (wsms *WebSocketManageService) SetWebSocketByToken(token string, webSocket *websocket.Conn)  {
//	wsms.mainSessionStorage.SetWebSocketToUserByToken(token, webSocket)
//}
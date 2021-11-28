package service

import (
	"backgammon/domain"
)

type WebSocketManageService struct {
	mainSessionStorage domain.SessionStorage
}

func NewWebSocketManageService(mainSessionStorage domain.SessionStorage) *WebSocketManageService {
	return &WebSocketManageService{mainSessionStorage: mainSessionStorage}
}

//func (wsms *WebSocketManageService) SetWebSocketByToken(token string, webSocket *websocket.Conn)  {
//	wsms.mainSessionStorage.SetWebSocketToUserByToken(token, webSocket)
//}

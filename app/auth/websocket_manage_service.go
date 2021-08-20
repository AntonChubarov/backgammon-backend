package auth

import "github.com/gorilla/websocket"

type WebSocketManageService struct {
	mainSessionStorage SessionStorage
}

func NewWebSocketManageService(mainSessionStorage SessionStorage) *WebSocketManageService {
	return &WebSocketManageService{mainSessionStorage: mainSessionStorage}
}

func (wsms *WebSocketManageService) CheckToken(token string) error {
	if !wsms.mainSessionStorage.IsTokenValid(token) {
		return ErrorInvalidToken
	}
	return nil
}

func (wsms *WebSocketManageService) SetWebSocketByToken(token string, webSocket *websocket.Conn)  {
	wsms.mainSessionStorage.SetWebSocketToUserByToken(token, webSocket)
}
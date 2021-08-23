package dal

import (
	"backgammon/app/auth"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type MainSessionStorage struct {
	storage map[string]*auth.UserSessionData
}

func NewMainSessionStorage() *MainSessionStorage {
	storage := make(map[string]*auth.UserSessionData)
	//storage := restoreFromDatabase()

	mainSessionStorage := &MainSessionStorage{
		storage: storage,
	}
	go mainSessionStorage.overdueTokenCleaner()
	return mainSessionStorage
}

func (mss *MainSessionStorage) AddNewUser(data *auth.UserSessionData) {
	data.ExpiryTime = time.Now().UTC().Add(30 * time.Second)

	mss.storage[data.Token] = data
}

func (mss *MainSessionStorage) UpdateTokenExpiryTime(token string) {
	temp := mss.storage[token]
	newExpiryTime := time.Now().UTC().Add(30 * time.Second)
	temp.ExpiryTime = newExpiryTime
	mss.storage[token] = temp
}

func (mss *MainSessionStorage) DeleteUserByToken(token string) {
	delete(mss.storage, token)
}

func (mss *MainSessionStorage) GetTokenByUUID(uuid string) (token string, wasFound bool) {
	for _, user := range mss.storage {
		if user.UserUUID == uuid {
			mss.UpdateTokenExpiryTime(user.Token)
			return user.Token, true
		}
	}
	return "", false
}

func (mss *MainSessionStorage) IsTokenValid(token string) bool {
	_, ok := mss.storage[token]
	return ok
}

func (mss *MainSessionStorage) SetWebSocketToUserByToken(token string, webSocket *websocket.Conn) {
	mss.storage[token].WebSocket = webSocket

	// for testing
	log.Println("to user", mss.storage[token].UserUUID, "set websocket", &webSocket)
}

func restoreFromDatabase() map[string]*auth.UserSessionData {
	panic("implement me")
}

func (mss *MainSessionStorage) overdueTokenCleaner() {
	for {
		time.Sleep(5 * time.Second)
		for _, user := range mss.storage {
			if user.ExpiryTime.Before(time.Now()) {
				user.WebSocket.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "you have been inactive for too long"))
				user.WebSocket.Close()
				delete(mss.storage, user.Token)
				log.Println("token", user.Token, "deleted")
			}
		}
	}
}
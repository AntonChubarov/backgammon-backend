package auth

import (
	"backgammon/app/auth"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type MainSessionStorage struct {
	storage map[string]*auth.UserSessionData
	mutex sync.RWMutex
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

	mss.mutex.Lock()
	mss.storage[data.Token] = data
	mss.mutex.Unlock()
}

func (mss *MainSessionStorage) UpdateTokenExpiryTime(token string) {
	mss.mutex.RLocker()
	temp := mss.storage[token]
	mss.mutex.RUnlock()

	newExpiryTime := time.Now().UTC().Add(30 * time.Second)
	temp.ExpiryTime = newExpiryTime

	mss.mutex.Lock()
	mss.storage[token] = temp
	mss.mutex.Unlock()
}

func (mss *MainSessionStorage) DeleteUserByToken(token string) {
	delete(mss.storage, token)
}

func (mss *MainSessionStorage) GetTokenByUUID(uuid string) (token string, wasFound bool) {
	mss.mutex.RLock()
	for _, user := range mss.storage {
		if user.UserUUID == uuid {
			mss.UpdateTokenExpiryTime(user.Token)
			return user.Token, true
		}
	}
	mss.mutex.RUnlock()

	return "", false
}

func (mss *MainSessionStorage) IsTokenValid(token string) bool {
	mss.mutex.RLock()
	_, ok := mss.storage[token]
	mss.mutex.RUnlock()

	return ok
}

func (mss *MainSessionStorage) SetWebSocketToUserByToken(token string, webSocket *websocket.Conn) {
	mss.mutex.RLock()
	mss.storage[token].WebSocket = webSocket

	// for testing
	log.Println("to user", mss.storage[token].UserUUID, "set websocket", &webSocket)
	mss.mutex.RUnlock()
}

func restoreFromDatabase() map[string]*auth.UserSessionData {
	panic("implement me")
}

// need to be redeveloped
func (mss *MainSessionStorage) overdueTokenCleaner() {
	
	for {
		time.Sleep(5 * time.Second)
		log.Println("Token cleaner start")
		mss.mutex.RLock()
		for key, _ := range mss.storage {
			if mss.storage[key].ExpiryTime.Before(time.Now()) {
				mss.storage[key].WebSocket.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "you have been inactive for too long"))
				mss.storage[key].WebSocket.Close()
				mss.mutex.RUnlock()
				mss.mutex.Lock()
				delete(mss.storage, mss.storage[key].Token)
				mss.mutex.Unlock()
				mss.mutex.RLock()
				log.Println("token", mss.storage[key].Token, "deleted")
			}
		}
		mss.mutex.RUnlock()
		log.Println("Token cleaner finish")
	}
}
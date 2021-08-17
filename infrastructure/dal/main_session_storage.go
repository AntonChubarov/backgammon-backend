package dal

import (
	"backgammon/domain/auth"
	"log"
	"time"
)

type MainSessionStorage struct {
	storage map[string]auth.UserSessionData
}

func NewMainSessionStorage() *MainSessionStorage {
	storage := make(map[string]auth.UserSessionData)
	//storage := restoreFromDatabase()

	mainSessionStorage := &MainSessionStorage{
		storage: storage,
	}
	go mainSessionStorage.overdueTokenCleaner()
	return mainSessionStorage
}

func (mss *MainSessionStorage) AddNewUser(data auth.UserSessionData) {
	data.ExpiryTime = time.Now().Add(time.Minute)
	mss.storage[data.Token] = data
	//log.Println(data)
}

func (mss *MainSessionStorage) UpdateTokenExpiryTime(token string) {
	temp := mss.storage[token]
	newExpiryTime := time.Now().Add(time.Minute)
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

func restoreFromDatabase() map[string]auth.UserSessionData {
	panic("implement me")
}

func (mss *MainSessionStorage) overdueTokenCleaner() {
	for {
		time.Sleep(5 * time.Second)
		for _, user := range mss.storage {
			if user.ExpiryTime.Before(time.Now()) {
				delete(mss.storage, user.Token)
				log.Println("token", user.Token, "deleted")
			}
		}
	}
}
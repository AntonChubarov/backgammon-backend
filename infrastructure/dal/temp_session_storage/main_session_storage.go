package temp_session_storage

import (
	"backgammon/app/auth"
	"backgammon/domain/authdomain"
	"log"
	"sync"
	"time"
)

type MainSessionStorage struct {
	storage map[authdomain.Token]authdomain.SessionData
	mutex sync.RWMutex
}

func NewMainSessionStorage() *MainSessionStorage {
	storage := make(map[authdomain.Token]authdomain.SessionData)
	//storage := restoreFromDatabase()

	mainSessionStorage := &MainSessionStorage{
		storage: storage,
	}
	go mainSessionStorage.overdueTokenCleaner()
	return mainSessionStorage
}

func (mss *MainSessionStorage) AddSession (data authdomain.SessionData) error {
	data.ExpiryTime = authdomain.ExpiryTime(time.Now().UTC().Add(30 * time.Second))

	mss.mutex.Lock()
	mss.storage[data.Token] = data
	mss.mutex.Unlock()

	return nil
}

func (mss *MainSessionStorage) UpdateSession(token authdomain.Token, data authdomain.SessionData) error {
	data.ExpiryTime = authdomain.ExpiryTime(time.Now().UTC().Add(30 * time.Second))

	mss.mutex.Lock()
	mss.storage[token] = data
	mss.mutex.Unlock()

	return nil
}

func (mss *MainSessionStorage) DeleteSession(token authdomain.Token) error {
	delete(mss.storage, token)

	return nil
}

func (mss *MainSessionStorage) GetSessionSByUUID(uuid authdomain.UUID) (authdomain.SessionData, error) {
	mss.mutex.RLock()
	for _, user := range mss.storage {
		if user.UUID == uuid {
			user.ExpiryTime = authdomain.ExpiryTime(time.Now().UTC().Add(30 * time.Second))
			return user, nil
		}
	}
	mss.mutex.RUnlock()

	return authdomain.SessionData{}, auth.ErrorNoActiveSessions
}

func (mss *MainSessionStorage) GetSessionByToken(token authdomain.Token) (authdomain.SessionData, error) {
	mss.mutex.RLock()
	session, ok := mss.storage[token]
	mss.mutex.RUnlock()

	if ok {
		return session, nil
	}
	return authdomain.SessionData{}, auth.ErrorInvalidToken
}



func restoreFromDatabase() map[string]authdomain.SessionData {
	panic("implement me")
}

// need to be redeveloped
func (mss *MainSessionStorage) overdueTokenCleaner() {
	
	for {
		time.Sleep(5 * time.Second)
		log.Println("Token cleaner start")
		mss.mutex.RLock()
		for key, _ := range mss.storage {
			if time.Time(mss.storage[key].ExpiryTime).Before(time.Now()) {
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
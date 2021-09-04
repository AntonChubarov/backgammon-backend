package ram_user_storage

import (
	"backgammon/app/auth"
	"backgammon/domain/authdomain"
	"github.com/viney-shih/go-lock"
)

type UserStorageRAM struct {
	storage map[authdomain.UUID]authdomain.UserData
	mu      lock.RWMutex //sync.RWMutex
}

func (u *UserStorageRAM) UpdateUser(uuid authdomain.UUID, data *authdomain.UserData) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	_, ok := u.storage[uuid]
	if !ok {
		return auth.ErrorUserNotRegistered
	}
	u.storage[uuid] = authdomain.UserData{
		UUID:     uuid,
		UserName: data.UserName,
		Password: data.Password,
	}
	return nil
}

func NewUserStorageRAM() *UserStorageRAM {
	return &UserStorageRAM{
		storage: make(map[authdomain.UUID]authdomain.UserData, 0),
		mu:      lock.NewCASMutex(), //sync.RWMutex{},
	}
}

func (u *UserStorageRAM) AddNewUser(data authdomain.UserData) error {

	if u.checkUser(&data) {
		return auth.ErrorUserExists
	}

	u.mu.Lock()
	defer u.mu.Unlock()
	u.storage[data.UUID] = data
	return nil
}
func (u *UserStorageRAM) checkUser(data *authdomain.UserData) bool {
	//u.RLock()
	//defer u.RUnlock()
	if _, err := u.GetUserByUUID(data.UUID); err == nil {
		return true
	}
	if _, err := u.GetUserByUsername(data.UserName); err == nil {
		return true
	}
	return false
}

func (u *UserStorageRAM) GetUserByUsername(name authdomain.UserName) (authdomain.UserData, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, v := range u.storage {
		if v.UserName == name {
			return v, nil
		}
	}
	return authdomain.UserData{}, auth.ErrorUserNotRegistered
}

func (u *UserStorageRAM) GetUserByUUID(uuid authdomain.UUID) (authdomain.UserData, error) {
	u.mu.RLock()
	v, ok := u.storage[uuid]
	u.mu.RUnlock()

	if !ok {
		return authdomain.UserData{}, auth.ErrorUserNotRegistered
	}
	return v, nil
}

func (u *UserStorageRAM) RemoveUser(uuid authdomain.UUID) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, ok := u.storage[uuid]; !ok {
		return auth.ErrorUserNotRegistered
	}
	delete(u.storage, uuid)
	return nil
}

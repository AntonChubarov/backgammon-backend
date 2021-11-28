package repository

import (
	"backgammon/app/service"
	"backgammon/domain"
	"github.com/viney-shih/go-lock"
)

type UserStorageRAM struct {
	storage map[domain.UUID]domain.UserData
	mu      lock.RWMutex //sync.RWMutex
}

func (u *UserStorageRAM) UpdateUser(uuid domain.UUID, data *domain.UserData) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	_, ok := u.storage[uuid]
	if !ok {
		return service.ErrorUserNotRegistered
	}
	u.storage[uuid] = domain.UserData{
		UUID:     uuid,
		UserName: data.UserName,
		Password: data.Password,
	}
	return nil
}

func NewUserStorageRAM() *UserStorageRAM {
	return &UserStorageRAM{
		storage: make(map[domain.UUID]domain.UserData, 0),
		mu:      lock.NewCASMutex(), //sync.RWMutex{},
	}
}

func (u *UserStorageRAM) AddNewUser(data domain.UserData) error {

	if u.checkUser(&data) {
		return service.ErrorUserExists
	}

	u.mu.Lock()
	defer u.mu.Unlock()
	u.storage[data.UUID] = data
	return nil
}
func (u *UserStorageRAM) checkUser(data *domain.UserData) bool {
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

func (u *UserStorageRAM) GetUserByUsername(name domain.UserName) (domain.UserData, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, v := range u.storage {
		if v.UserName == name {
			return v, nil
		}
	}
	return domain.UserData{}, service.ErrorUserNotRegistered
}

func (u *UserStorageRAM) GetUserByUUID(uuid domain.UUID) (domain.UserData, error) {
	u.mu.RLock()
	v, ok := u.storage[uuid]
	u.mu.RUnlock()

	if !ok {
		return domain.UserData{}, service.ErrorUserNotRegistered
	}
	return v, nil
}

func (u *UserStorageRAM) RemoveUser(uuid domain.UUID) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, ok := u.storage[uuid]; !ok {
		return service.ErrorUserNotRegistered
	}
	delete(u.storage, uuid)
	return nil
}

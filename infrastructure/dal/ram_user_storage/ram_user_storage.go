package ram_user_storage

import (
	"backgammon/app/auth"
	"backgammon/domain/authdomain"
	"sync"
)

type UserStorageRAM struct {
	storage map[authdomain.UUID]*authdomain.UserData
	sync.RWMutex
}

func (u *UserStorageRAM) UpdateUser(uuid authdomain.UUID, data *authdomain.UserData) error {
	u.RLock()
	defer u.RUnlock()
	user, ok:= u.storage[uuid]
	if !ok {return auth.ErrorUserNotRegistered}
	user.Update(data)
	return nil
}

func NewUserStorageRAM() *UserStorageRAM {
	return &UserStorageRAM{
		storage: make(map[authdomain.UUID]*authdomain.UserData,0),
	}
}

func (u *UserStorageRAM) AddNewUser(data *authdomain.UserData) error {


	if data==nil {return auth.ErrorNullArgument}
	if u.checkUser(data) {return auth.ErrorUserExists}

	u.Lock()
	defer u.Unlock()
	u.storage[data.UUID]=data
	return nil
}
func (u *UserStorageRAM) checkUser (data *authdomain.UserData) bool {
	//u.RLock()
	//defer u.RUnlock()
	if user, _:=u.GetUserByUUID(data.UUID); user!=nil {
		return true
	}
	if user, _:=u.GetUserByUsername(data.UserName); user!=nil {
		return true
	}
	return false
}

func (u *UserStorageRAM) GetUserByUsername(name authdomain.UserName) (*authdomain.UserData, error) {
	u.RLock()
	defer u.RUnlock()
	for _ , v :=range u.storage {
		if v.UserName==name {
			return v, nil
		}
	}
	return nil, auth.ErrorUserNotRegistered
}

func (u *UserStorageRAM) GetUserByUUID(uuid authdomain.UUID) (*authdomain.UserData, error) {
	u.RLock()
	v, ok :=u.storage[uuid]
	u.RUnlock()

	if !ok {
		return nil, auth.ErrorUserNotRegistered
	}
	return v, nil
}

func (u *UserStorageRAM) RemoveUser(uuid authdomain.UUID) error {
	u.Lock()
	defer u.Unlock()
	if _, ok :=u.storage[uuid]; !ok {
		return auth.ErrorUserNotRegistered
	}
	delete(u.storage,uuid)
	return nil
}
package repository

import (
	"backgammon/app/service"
	"backgammon/domain"
	"backgammon/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestUserStorageRAM_AddNewUser_single(t *testing.T) {
	var storage domain.UserStorage
	storage = NewUserStorageRAM()

	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(gofakeit.Password(true, true, true, false, false, 10)),
	}
	err := storage.AddNewUser(userData)
	assert.Equal(t, nil, err)

	readUser, err := storage.GetUserByUUID(userData.UUID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, readUser)
	assert.Equal(t, userData, readUser)
}

func TestUserStorageRAM_AddNewUser_duplicate(t *testing.T) {

	var storage domain.UserStorage
	storage = NewUserStorageRAM()

	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(gofakeit.Password(true, true, true, false, false, 10)),
	}
	err := storage.AddNewUser(userData)
	err = storage.AddNewUser(userData)
	assert.Equal(t, service.ErrorUserExists, err)
}

func TestUserStorageRAM_AddNewUser_NonConcurrent(t *testing.T) {
	var storage domain.UserStorage
	storage = NewUserStorageRAM()
	count := 1000

	for i := 1; i <= count; i++ {
		userData := domain.UserData{
			UUID:     utils.GenerateUUID(),
			UserName: domain.UserName(gofakeit.Password(true, true, true, false, false, 32)),
			Password: domain.Password(gofakeit.Password(true, true, true, false, false, 10)),
		}
		err := storage.AddNewUser(userData)
		assert.Equal(t, nil, err)
	}

}

func TestUserStorageRAM_AddNewUser_Concurrent(t *testing.T) {
	count := 1000
	sl1 := makeRamUserArray(count)
	sl2 := makeRamUserArray(count)
	sl3 := makeRamUserArray(count)
	sl4 := makeRamUserArray(count)
	storage := NewUserStorageRAM()
	wg := sync.WaitGroup{}
	wg.Add(4)
	rt := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			assert.Nil(t, storage.AddNewUser(sl[i]))
		}
	}
	go rt(sl1)
	go rt(sl2)
	go rt(sl3)
	go rt(sl4)
	wg.Wait()
	assert.True(t, true)

}

func TestUserStorageRAM_AddNewUser_Concurrent_verify(t *testing.T) {
	count := 1000
	sl1 := makeRamUserArray(count)
	sl2 := makeRamUserArray(count)
	sl3 := makeRamUserArray(count)
	sl4 := makeRamUserArray(count)
	storage := NewUserStorageRAM()
	wg := sync.WaitGroup{}
	wg.Add(4)

	rt := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			err := storage.AddNewUser(sl[i])
			assert.Nil(t, err)
		}
	}
	go rt(sl1)
	go rt(sl2)
	go rt(sl3)
	go rt(sl4)

	wg.Wait()

	sl := append(sl1, sl2...)
	sl = append(sl, sl3...)
	sl = append(sl, sl4...)

	for i := range sl {
		user, err := storage.GetUserByUUID(sl[i].UUID)
		assert.Nil(t, err)
		assert.Equal(t, sl[i], user)
	}
}

func TestUserStorageRAM_ConcurrentRandomAccess(t *testing.T) {
	count := 1000
	storage := NewUserStorageRAM()
	sl1 := makeRamUserArray(count)
	sl2 := makeRamUserArray(count)
	sl3 := makeRamUserArray(count)
	wg := sync.WaitGroup{}
	wg.Add(2)
	rt := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			err := storage.AddNewUser(sl[i])
			assert.Nil(t, err)
		}
	}
	go rt(sl1)
	go rt(sl2)
	wg.Wait()

	rName := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			user, err := storage.GetUserByUUID(sl[i].UUID)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], user)
		}
	}

	rUuid := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			user, err := storage.GetUserByUsername(sl[i].UserName)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], user)
		}
	}
	wg.Add(3)
	go rName(sl1)
	go rUuid(sl2)
	go rt(sl3)
	wg.Wait()
	assert.True(t, true)

}

func TestUserStorageRAM_FullRandomAccess(t *testing.T) {
	count := 10000
	var readsCounter int64
	maxInterval := 500
	storage := NewUserStorageRAM()
	sl1 := makeRamUserArray(count)
	sl2 := makeRamUserArray(count)
	sl3 := makeRamUserArray(count)
	sl4 := makeRamUserArray(count)
	sl5 := makeRamUserArray(count)
	sl6 := makeRamUserArray(count)

	wg := sync.WaitGroup{}
	wg.Add(4)
	rt := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			err := storage.AddNewUser(sl[i])
			assert.Nil(t, err)
		}
	}
	go rt(sl1)
	go rt(sl2)
	go rt(sl3)
	go rt(sl4)

	wg.Wait()

	rUuid := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			user, err := storage.GetUserByUUID(sl[i].UUID)
			atomic.AddInt64(&readsCounter, 1)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], user)
		}
	}

	rName := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			user, err := storage.GetUserByUsername(sl[i].UserName)
			atomic.AddInt64(&readsCounter, 1)
			assert.Nil(t, err)
			assert.Equal(t, sl[i], user)
		}
	}

	rDel := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			tmp := atomic.LoadInt64(&readsCounter)
			atomic.StoreInt64(&readsCounter, 0)
			assert.Less(t, tmp, int64(maxInterval))
			err := storage.RemoveUser(sl[i].UUID)
			assert.Nil(t, err)
		}
	}

	rUpd := func(slOrg []domain.UserData, slUpd []domain.UserData) {
		defer wg.Done()
		for i := range slOrg {
			tmp := atomic.LoadInt64(&readsCounter)
			atomic.StoreInt64(&readsCounter, 0)
			assert.Less(t, tmp, int64(maxInterval))
			storage.UpdateUser(slOrg[i].UUID, &slUpd[i])
			expected := slUpd[i].UserName
			u, err := storage.GetUserByUUID(slOrg[i].UUID)
			assert.Nil(t, err)
			actual := u.UserName
			assert.Equal(t, expected, actual)
		}
	}
	wg.Add(5)
	go rUuid(sl1)
	go rName(sl2)
	go rDel(sl3)
	go rUpd(sl4, sl5)
	go rt(sl6)

	for i := 1; i <= 100; i++ {
		time.Sleep(500 * time.Microsecond)
		wg.Add(1)
		go rName(sl1)

		wg.Add(1)
		go rUuid(sl1)
	}

	wg.Wait()
	assert.True(t, true)
}

func TestUserStorageRAM_GetUserByUUID_single(t *testing.T) {
	var storage domain.UserStorage
	storage = NewUserStorageRAM()

	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(gofakeit.Password(true, true, true, false, false, 10)),
	}
	err := storage.AddNewUser(userData)
	assert.Equal(t, nil, err)

	readUser, err2 := storage.GetUserByUUID(userData.UUID)
	assert.Nil(t, err2)
	assert.Equal(t, userData, readUser)
}

func makeRamUserArray(count int) []domain.UserData {
	sl := make([]domain.UserData, 0, count)
	for i := 0; i < count; i++ {
		userData := domain.UserData{
			UUID:     utils.GenerateUUID(),
			UserName: domain.UserName(gofakeit.Password(true, true, true, false, false, 32)),
			Password: domain.Password(gofakeit.Password(true, true, true, false, false, 10)),
		}
		sl = append(sl, userData)
	}
	return sl
}
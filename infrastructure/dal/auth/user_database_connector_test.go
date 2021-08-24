package auth

import (
	"backgammon/app/auth"
	"backgammon/config"
	"backgammon/domain/authdomain"
	"backgammon/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var ServerConfig = config.ServerConfig{
	Database: config.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "backgammonadmin",
		Password: "backgammon",
		Name:     "backgammon",
	},
}

func TestUserStorageRAM_AddNewUser_single(t *testing.T) {
	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)

	userData:=authdomain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: authdomain.UserName(gofakeit.Username()),
		Password: authdomain.Password(gofakeit.Password(true, true, true, false, false, 10)),
	}
	err:=storage.AddNewUser(&userData)
	assert.Equal(t, nil, err)

	readUser, err:= storage.GetUserByUUID(userData.UUID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, readUser)
	assert.Equal(t, userData, *readUser)
}

func TestUserStorageRAM_AddNewUser_duplicate(t *testing.T) {

	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)

	userData:=authdomain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: authdomain.UserName(gofakeit.Username()),
		Password: authdomain.Password(gofakeit.Password(true, true, true, false, false, 10)),
	}
	err:=storage.AddNewUser(&userData)
	err=storage.AddNewUser(&userData)
	assert.Equal(t, auth.ErrorUserExists, err)
}

func TestUserStorageRAM_AddNewUser_NonConcurrent(t *testing.T) {
	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)
	count:=1000

	for i:=1; i<=count; i++ {
		userData:=authdomain.UserData{
			UUID:     utils.GenerateUUID(),
			UserName: authdomain.UserName(gofakeit.Password(true, true, true, false, false, 32)),
			Password: authdomain.Password(gofakeit.Password(true, true, true, false, false, 10)),
		}
		err:=storage.AddNewUser(&userData)
		assert.Equal(t, nil, err)
	}

}

func TestUserStorageRAM_AddNewUser_Concurrent(t *testing.T) {
	var storage authdomain.UserStorage
	count:=1000
	sl1:=makeUserArray(count)
	sl2:=makeUserArray(count)
	sl3:=makeUserArray(count)
	sl4:=makeUserArray(count)
	storage = NewDatabaseConnector(&ServerConfig)
	wg:=sync.WaitGroup{}
	wg.Add(4)
	rt:= func(sl []authdomain.UserData) {
		defer wg.Done()
		for i:=range sl {
			assert.Nil(t, storage.AddNewUser(&sl[i]))
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
	count:=1000
	sl1:=makeUserArray(count)
	sl2:=makeUserArray(count)
	sl3:=makeUserArray(count)
	sl4:=makeUserArray(count)
	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)
	var wg sync.WaitGroup
	wg.Add(4)

	rt:= func(sl []authdomain.UserData) {
		defer wg.Done()
		for i:=range sl {
			err:=storage.AddNewUser(&sl[i])
			assert.Nil(t, err)
		}
	}
	go rt(sl1)
	go rt(sl2)
	go rt(sl3)
	go rt(sl4)

	wg.Wait()

	sl:=append(sl1, sl2...)
	sl=append(sl, sl3...)
	sl=append(sl, sl4...)



	for i:=range sl {
		user, err:= storage.GetUserByUUID(sl[i].UUID)
		assert.Nil(t, err)
		assert.Equal (t, sl[i],  *user )
	}
}

func TestUserStorageRAM_ConcurrentRandomAccess (t *testing.T) {
	count:=1000
	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)
	sl1:=makeUserArray(count)
	sl2:=makeUserArray(count)
	sl3:=makeUserArray(count)
	wg:=sync.WaitGroup{}
	wg.Add(2)
	rt:= func(sl []authdomain.UserData) {
		defer wg.Done()
		for i:=range sl {
			err:=storage.AddNewUser(&sl[i])
			assert.Nil(t, err)
		}
	}
	go rt(sl1)
	go rt(sl2)
	wg.Wait()

	rName:= func (sl []authdomain.UserData){
		defer wg.Done()
		for i:=range sl {
			user, err:= storage.GetUserByUUID(sl[i].UUID)
			assert.Nil(t, err)
			assert.Equal (t, sl[i],  *user )
		}
	}

	rUuid:=func (sl []authdomain.UserData){
		defer wg.Done()
		for i:=range sl {
			user, err:= storage.GetUserByUsername(sl[i].UserName)
			assert.Nil(t, err)
			assert.Equal (t, sl[i],  *user )
		}
	}
	wg.Add(3)
	go rName(sl1)
	go rUuid(sl2)
	go rt(sl3)
	wg.Wait()
	assert.True(t, true)




}

func TestUserStorageRAM_FullRandomAccess (t *testing.T) {
	count:=1000
	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)
	sl1:=makeUserArray(count)
	sl2:=makeUserArray(count)
	sl3:=makeUserArray(count)
	sl4:=makeUserArray(count)
	sl5:=makeUserArray(count)

	wg:=sync.WaitGroup{}
	wg.Add(3)
	rt:= func(sl []authdomain.UserData) {
		defer wg.Done()
		for i:=range sl {
			err:=storage.AddNewUser(&sl[i])
			assert.Nil(t, err)
		}
	}
	go rt(sl1)
	go rt(sl2)
	go rt(sl3)


	wg.Wait()

	rName:= func (sl []authdomain.UserData){
		defer wg.Done()
		for i:=range sl {
			user, err:= storage.GetUserByUUID(sl[i].UUID)
			assert.Nil(t, err)
			assert.Equal (t, sl[i],  *user )
		}
	}

	rUuid:=func (sl []authdomain.UserData){
		defer wg.Done()
		for i:=range sl {
			user, err:= storage.GetUserByUsername(sl[i].UserName)
			assert.Nil(t, err)
			assert.Equal (t, sl[i],  *user )
		}
	}

	rDel:=func (sl []authdomain.UserData) {
		defer wg.Done()
		for i:=range sl {
			err := storage.RemoveUser(sl[i].UUID)
			assert.Nil(t, err)
		}
	}

	rUpd:=func (slOrg []authdomain.UserData, slUpd []authdomain.UserData) {
		defer wg.Done()
		for i:=range slOrg {
			storage.UpdateUser(slOrg[i].UUID, &slUpd[i])
			expected:=slUpd[i].UserName
			u, err:=storage.GetUserByUUID(slOrg[i].UUID)
			assert.Nil(t, err)
			actual:=u.UserName
			assert.Equal(t, expected, actual)

		}
	}
	wg.Add(5)
	go rName(sl1)
	go rUuid(sl2)
	go rDel(sl3)
	go rUpd(sl1,sl4)
	go rt(sl5)
	wg.Wait()
	assert.True(t, true)
}


func TestUserStorageRAM_GetUserByUUID_single(t *testing.T) {
	var storage authdomain.UserStorage
	storage = NewDatabaseConnector(&ServerConfig)

	userData:=authdomain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: authdomain.UserName(gofakeit.Username()),
		Password: authdomain.Password(gofakeit.Password(true, true, true, false, false, 10)),
	}
	err:=storage.AddNewUser(&userData)
	assert.Equal(t, nil, err)

	readUser, err2:=storage.GetUserByUUID(userData.UUID)
	assert.Nil(t, err2)
	assert.Equal(t, userData, *readUser)

}

func makeUserArray(count int) []authdomain.UserData {
	sl:=make([]authdomain.UserData, 0, count)
	for i:=0; i<count; i++ {
		userData:=authdomain.UserData{
			UUID:     utils.GenerateUUID(),
			UserName: authdomain.UserName(gofakeit.Password(true, true, true, false, false, 32)),
			Password: authdomain.Password(gofakeit.Password(true, true, true, false, false, 10)),
		}

		sl=append(sl, userData)
	}
	return sl
}


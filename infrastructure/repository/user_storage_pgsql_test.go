package repository

import (
	"backgammon/app/service"
	"backgammon/config"
	"backgammon/domain"
	"backgammon/utils"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"testing"
)

type UserDBDTOPass struct {
	UUID         string `db:"useruuid"`
	Username     string `db:"username"`
	PasswordHash string `db:"userpassword"`
	Pass         int    `db:"pass"`
}

var hasher = service.NewHasherSHA256()

var ServerConfig = config.ServerConfig{
	Database: config.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "backgammonadmin",
		Password: "backgammon",
		Name:     "backgammon",
	},
}

func TestUserStoragePGSQL_AddNewUser_single(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

	password, _ := hasher.HashString(gofakeit.Password(true, true, true, false, false, 10))
	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(password),
	}
	err := storage.AddNewUser(userData)
	assert.Equal(t, nil, err)

	readUser, err := storage.GetUserByUUID(userData.UUID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, readUser)
	assert.Equal(t, userData, readUser)
}

func TestUserStoragePGSQL_AddNewUser_duplicate(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

	password, _ := hasher.HashString(gofakeit.Password(true, true, true, false, false, 10))
	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(password),
	}
	err := storage.AddNewUser(userData)
	err2 := storage.AddNewUser(userData)
	assert.Nil(t, err)
	assert.NotNil(t, err2)
}

func TestUserStoragePGSQL_AddNewUser_NonConcurrent(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

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

func TestUserStoragePGSQL_AddNewUser_Concurrent(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

	count := 1000
	sl1 := makePgsqlUserArray(count)
	sl2 := makePgsqlUserArray(count)
	sl3 := makePgsqlUserArray(count)
	sl4 := makePgsqlUserArray(count)

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

func TestUserStoragePGSQL_AddNewUser_Concurrent_verify(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

	count := 1000
	sl1 := makePgsqlUserArray(count)
	sl2 := makePgsqlUserArray(count)
	sl3 := makePgsqlUserArray(count)
	sl4 := makePgsqlUserArray(count)

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

func TestUserStoragePGSQL_ConcurrentRandomAccess(t *testing.T) {
	count := 1000
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor
	sl1 := makePgsqlUserArray(count)
	sl2 := makePgsqlUserArray(count)
	sl3 := makePgsqlUserArray(count)
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

func TestUserDataStoragePGSQL_RemoveUser_single(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

	password, _ := hasher.HashString(gofakeit.Password(true, true, true, false, false, 10))
	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(password),
	}
	err := storage.AddNewUser(userData)
	assert.Equal(t, nil, err)

	readUser, err := storage.GetUserByUUID(userData.UUID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, readUser)
	assert.Equal(t, userData, readUser)

	err2 := storage.RemoveUser(userData.UUID)
	assert.Nil(t, err2)
	readUser2, err3 := storage.GetUserByUUID(userData.UUID)
	assert.Equal(t, domain.UserData{}, readUser2)
	assert.Equal(t, service.ErrorUserNotRegistered, err3)

}

func TestUserDataStoragePGSQL_UpdateUser_single(t *testing.T) {
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor

	password, _ := hasher.HashString(gofakeit.Password(true, true, true, false, false, 10))
	userData := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(password),
	}

	password2, _ := hasher.HashString(gofakeit.Password(true, true, true, false, false, 10))
	userData2 := domain.UserData{
		UUID:     utils.GenerateUUID(),
		UserName: domain.UserName(gofakeit.Username()),
		Password: domain.Password(password2),
	}

	err := storage.AddNewUser(userData)
	assert.Equal(t, nil, err)

	readUser, err := storage.GetUserByUUID(userData.UUID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, readUser)
	assert.Equal(t, userData, readUser)

	err2 := storage.UpdateUser(userData.UUID, &userData2)
	assert.Nil(t, err2)
	readUser2, err3 := storage.GetUserByUUID(userData.UUID)
	userData2.UUID = userData.UUID
	assert.Nil(t, err3)
	assert.Equal(t, userData2, readUser2)
}

func TestUserStoragePGSQL_FullRandomAccess_prepare(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ServerConfig.Database.Host,
		ServerConfig.Database.Port,
		ServerConfig.Database.User,
		ServerConfig.Database.Password,
		ServerConfig.Database.Name,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.Exec("DELETE FROM users_temp")

	count := 10000
	sl1 := makePgsqlUserArray(count)
	sl2 := makePgsqlUserArray(count)
	sl3 := makePgsqlUserArray(count)
	sl4 := makePgsqlUserArray(count)
	sl5 := makePgsqlUserArray(count)

	saveUserArray(1, sl1)
	saveUserArray(2, sl2)
	saveUserArray(3, sl3)
	saveUserArray(4, sl4)
	saveUserArray(5, sl5)
	assert.True(t, true)
}

func TestUserStoragePGSQL_FullRandomAccess(t *testing.T) {
	count := 10000
	stor := NewUserDataStoragePGSQL(&ServerConfig)
	defer stor.CloseDatabaseConnection()
	var storage domain.UserStorage
	storage = stor
	sl1 := makePgsqlUserArray(count)
	sl2 := makePgsqlUserArray(count)
	sl3 := makePgsqlUserArray(count)
	sl4 := makePgsqlUserArray(count)
	sl5 := makePgsqlUserArray(count)
	sl6 := makePgsqlUserArray(count)

	wg := sync.WaitGroup{}

	rt := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			err := storage.AddNewUser(sl[i])
			assert.Nil(t, err)
		}
	}
	wg.Add(4)
	go rt(sl1)
	go rt(sl2)
	go rt(sl3)
	go rt(sl4)
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

	rDel := func(sl []domain.UserData) {
		defer wg.Done()
		for i := range sl {
			err := storage.RemoveUser(sl[i].UUID)
			assert.Nil(t, err)
		}
	}

	rUpd := func(slOrg []domain.UserData, slUpd []domain.UserData) {
		defer wg.Done()
		for i := range slOrg {
			storage.UpdateUser(slOrg[i].UUID, &slUpd[i])
			expected := slUpd[i].UserName
			u, err := storage.GetUserByUUID(slOrg[i].UUID)
			assert.Nil(t, err)
			actual := u.UserName
			assert.Equal(t, expected, actual)

		}
	}
	wg.Add(5)
	go rName(sl1)
	go rUuid(sl2)
	go rDel(sl3)
	go rUpd(sl4, sl5)
	go rt(sl6)
	wg.Wait()
	assert.True(t, true)
}

func TestUserStoragePGSQL_GetUserByUUID_single(t *testing.T) {
	var storage domain.UserStorage
	storage = NewUserDataStoragePGSQL(&ServerConfig)

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

func loadUserArray(pass int) []domain.UserData {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ServerConfig.Database.Host,
		ServerConfig.Database.Port,
		ServerConfig.Database.User,
		ServerConfig.Database.Password,
		ServerConfig.Database.Name,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	defer db.Close()

	var users []UserDBDTO
	usersOut := make([]domain.UserData, 0, 0)

	err2 := db.Select(&users, "select username, userpassword, useruuid from users_temp where pass = $1", pass)
	if err != nil {
		panic(err2)
	}

	for i := range users {
		usersOut = append(usersOut, userDBDTOToUserData(users[i]))
	}
	return usersOut
}

func saveUserArray(pass int, sl []domain.UserData) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ServerConfig.Database.Host,
		ServerConfig.Database.Port,
		ServerConfig.Database.User,
		ServerConfig.Database.Password,
		ServerConfig.Database.Name,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	for i := range sl {
		userDTO := UserDBDTOPass{
			UUID:         string(sl[i].UUID),
			Username:     string(sl[i].UserName),
			PasswordHash: string(sl[i].Password),
			Pass:         pass,
		}
		_, err2 := db.NamedExec("insert into users_temp (useruuid, username, userpassword, pass) values (:useruuid, :username, :userpassword, :pass)",
			userDTO)
		if err2 != nil {
			log.Println(err2)
		}
	}
}

func makePgsqlUserArray(count int) []domain.UserData {
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

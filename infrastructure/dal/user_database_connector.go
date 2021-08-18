package dal

import (
	auth2 "backgammon/app/auth"
	"backgammon/config"
	"backgammon/domain/auth"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // shadow import
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // shadow import
	"log"
)

type DatabaseConnector struct {
	Database *sqlx.DB
}

func NewDatabaseConnector(config *config.ServerConfig) *DatabaseConnector {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &DatabaseConnector{
		Database: db,
	}
}

func (d *DatabaseConnector) CloseDatabaseConnection() {
	err := d.Database.Close()
	if err != nil {
		log.Println(err)
	}
}

func (d *DatabaseConnector) AddNewUser(data auth.UserAuthData) error {
	userDTO := UserDataToUserDBDTO(data)

	_, err := d.Database.NamedExec("insert into users (useruuid, userlogin, userpassword) values (:useruuid, :userlogin, :userpassword)",
		userDTO)
	if err != nil {
		log.Println("In dal.AddNewUser", err)
		return err
	}

	return nil
}

func (d *DatabaseConnector) IsUserExist(username string) (bool, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select userlogin, userpassword from users where userlogin = $1", username)
	if err != nil {
		log.Println("In dal.IsUserExist", err)
		return false, err
	}
	if users != nil {
		return true, nil
	}
	return false, nil
}

func (d *DatabaseConnector) GetUserByUsername(username string) (auth.UserAuthData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select userlogin, userpassword, useruuid from users where userlogin = $1", username)
	if err != nil {
		log.Println("In dal.GetUserByUsername", err)
		return auth.UserAuthData{}, err
	}
	if users == nil {
		return auth.UserAuthData{}, ErrorNoUserInDatabase
	}
	if len(users) == 1 {
		return UserDBDTOToUserData(users[0]), nil
	}
	if len(users) > 1 {
		return auth.UserAuthData{}, ErrorMoreThanOneUsernameRecord
	}
	return auth.UserAuthData{}, auth2.ErrorInvalidUsername
}

func (d *DatabaseConnector) UpdateUser(oldData, newData auth.UserAuthData) error {
	panic("implement me")
}

func (d *DatabaseConnector) RemoveUser(data auth.UserAuthData) error {
	panic("implement me")
}

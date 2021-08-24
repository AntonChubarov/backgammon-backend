package auth

import (
	auth2 "backgammon/app/auth"
	"backgammon/config"
	"backgammon/domain/authdomain"
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

func (d *DatabaseConnector) AddNewUser(data *authdomain.UserData) error {
	userDTO := UserDataToUserDBDTO(*data)

	_, err := d.Database.NamedExec("insert into users (useruuid, username, userpassword) values (:useruuid, :username, :userpassword)",
		userDTO)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}

	return nil
}

func (d *DatabaseConnector) GetUserByUsername(username authdomain.UserName) (*authdomain.UserData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select username, userpassword, useruuid from users where username = $1", username)
	if err != nil {
		//log.Println("In dal.GetUserByUsername", err)
		return &authdomain.UserData{}, err
	}
	if users == nil {
		return &authdomain.UserData{}, ErrorNoUserInDatabase
	}
	if len(users) == 1 {
		user := UserDBDTOToUserData(users[0])
		return &user, nil
	}
	if len(users) > 1 {
		return &authdomain.UserData{}, ErrorMoreThanOneUsernameRecord
	}
	return &authdomain.UserData{}, auth2.ErrorUserNotRegistered
}

func (d *DatabaseConnector) GetUserByUUID(uuid authdomain.UUID) (*authdomain.UserData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select username, userpassword, useruuid from users where useruuid = $1", string(uuid))
	if err != nil {
		//log.Println("In dal.GetUserByUsername", err)
		return &authdomain.UserData{}, err
	}
	if users == nil {
		return &authdomain.UserData{}, ErrorNoUserInDatabase
	}
	if len(users) == 1 {
		user := UserDBDTOToUserData(users[0])
		return &user, nil
	}
	if len(users) > 1 {
		return &authdomain.UserData{}, ErrorMoreThanOneUsernameRecord
	}
	return &authdomain.UserData{}, auth2.ErrorUserNotRegistered
}

func (d *DatabaseConnector) UpdateUser(uuid authdomain.UUID, data *authdomain.UserData) error {
	userDTO := UserDataToUserDBDTO(*data)

	_, err := d.Database.NamedExec(`UPDATE users SET username=:username, userpassword=:userpassword WHERE useruuid=:useruuid`,
		userDTO)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}
	return nil
}

func (d *DatabaseConnector) RemoveUser(uuid authdomain.UUID) error {
	_, err := d.Database.Exec("DELETE FROM users WHERE useruuid=$1", uuid)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}
	return nil
}
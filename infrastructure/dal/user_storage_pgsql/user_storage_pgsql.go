package user_storage_pgsql

import (
	"backgammon/app/auth"
	"backgammon/config"
	"backgammon/domain/authdomain"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // shadow import
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // shadow import
	"log"
)

var ErrorMoreThanOneUsernameRecord = fmt.Errorf("more than one user with this username, report to developers was automatically send")

type UserDBDTO struct {
	UUID         string `db:"useruuid"`
	Username     string `db:"username"`
	PasswordHash string `db:"userpassword"`
}

func userDataToUserDBDTO(user authdomain.UserData) UserDBDTO {
	return UserDBDTO{
		UUID:         string(user.UUID),
		Username:     string(user.UserName),
		PasswordHash: string(user.Password),
	}
}

func userDBDTOToUserData(user UserDBDTO) authdomain.UserData {
	return authdomain.UserData{
		UUID:     authdomain.UUID(user.UUID),
		UserName: authdomain.UserName(user.Username),
		Password: authdomain.Password(user.PasswordHash),
	}
}

type UserDataStoragePGSQL struct {
	Database *sqlx.DB
}

func NewUserDataStoragePGSQL(config *config.ServerConfig) *UserDataStoragePGSQL {
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

	return &UserDataStoragePGSQL{
		Database: db,
	}
}

func (d *UserDataStoragePGSQL) CloseDatabaseConnection() {
	err := d.Database.Close()
	if err != nil {
		log.Println(err)
	}
}

func (d *UserDataStoragePGSQL) AddNewUser(data *authdomain.UserData) error {
	userDTO := userDataToUserDBDTO(*data)

	_, err := d.Database.NamedExec("insert into users (useruuid, username, userpassword) values (:useruuid, :username, :userpassword)",
		userDTO)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}

	return nil
}

func (d *UserDataStoragePGSQL) GetUserByUsername(username authdomain.UserName) (authdomain.UserData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select username, userpassword, useruuid from users where username = $1", username)
	if err != nil {
		//log.Println("In dal.GetUserByUsername", err)
		return authdomain.UserData{}, err
	}
	if users == nil {
		return authdomain.UserData{}, auth.ErrorUserNotRegistered
	}
	if len(users) == 1 {
		user := userDBDTOToUserData(users[0])
		return user, nil
	}
	if len(users) > 1 {
		return authdomain.UserData{}, ErrorMoreThanOneUsernameRecord
	}
	return authdomain.UserData{}, auth.ErrorUserNotRegistered
}

func (d *UserDataStoragePGSQL) GetUserByUUID(uuid authdomain.UUID) (authdomain.UserData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select username, userpassword, useruuid from users where useruuid = $1", string(uuid))
	if err != nil {
		//log.Println("In dal.GetUserByUsername", err)
		return authdomain.UserData{}, err
	}
	if users == nil {
		return authdomain.UserData{}, auth.ErrorUserNotRegistered
	}
	if len(users) == 1 {
		user := userDBDTOToUserData(users[0])
		return user, nil
	}
	if len(users) > 1 {
		return authdomain.UserData{}, ErrorMoreThanOneUsernameRecord
	}
	return authdomain.UserData{}, auth.ErrorUserNotRegistered
}

func (d *UserDataStoragePGSQL) UpdateUser(uuid authdomain.UUID, data *authdomain.UserData) error {
	userDTO := userDataToUserDBDTO(*data)
	userDTO.UUID=string(uuid)

	_, err := d.Database.NamedExec(`UPDATE users SET username=:username, userpassword=:userpassword WHERE useruuid=:useruuid`,
		userDTO)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}
	return nil
}

func (d *UserDataStoragePGSQL) RemoveUser(uuid authdomain.UUID) error {
	_, err := d.Database.Exec("DELETE FROM users WHERE useruuid=$1", uuid)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}
	return nil
}
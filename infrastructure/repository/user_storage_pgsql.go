package repository

import (
	"backgammon/app/service"
	"backgammon/config"
	"backgammon/domain"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // shadow import
	"github.com/jbrodriguez/mlog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // shadow import
)

var ErrorMoreThanOneUsernameRecord = fmt.Errorf("more than one user with this username, report to developers was automatically send")

type UserDBDTO struct {
	UUID         string `db:"useruuid"`
	Username     string `db:"username"`
	PasswordHash string `db:"userpassword"`
}

func userDataToUserDBDTO(user domain.UserData) UserDBDTO {
	return UserDBDTO{
		UUID:         string(user.UUID),
		Username:     string(user.UserName),
		PasswordHash: string(user.Password),
	}
}

func userDBDTOToUserData(user UserDBDTO) domain.UserData {
	return domain.UserData{
		UUID:     domain.UUID(user.UUID),
		UserName: domain.UserName(user.Username),
		Password: domain.Password(user.PasswordHash),
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
		mlog.Fatalf("%s", err)
	}

	err = db.Ping()
	if err != nil {
		mlog.Fatalf("%s", err)
	}

	return &UserDataStoragePGSQL{
		Database: db,
	}
}

func (d *UserDataStoragePGSQL) CloseDatabaseConnection() {
	err := d.Database.Close()
	if err != nil {
		mlog.Warning("%s", err)
	}
}

func (d *UserDataStoragePGSQL) AddNewUser(data domain.UserData) error {
	userDTO := userDataToUserDBDTO(data)

	_, err := d.Database.NamedExec("insert into users (useruuid, username, userpassword) values (:useruuid, :username, :userpassword)",
		userDTO)
	if err != nil {
		return err
	}

	return nil
}

func (d *UserDataStoragePGSQL) GetUserByUsername(username domain.UserName) (domain.UserData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select username, userpassword, useruuid from users where username = $1", username)
	if err != nil {
		return domain.UserData{}, err
	}
	if users == nil {
		return domain.UserData{}, service.ErrorUserNotRegistered
	}
	if len(users) == 1 {
		user := userDBDTOToUserData(users[0])
		return user, nil
	}
	if len(users) > 1 {
		return domain.UserData{}, ErrorMoreThanOneUsernameRecord
	}
	return domain.UserData{}, service.ErrorUserNotRegistered
}

func (d *UserDataStoragePGSQL) GetUserByUUID(uuid domain.UUID) (domain.UserData, error) {
	var users []UserDBDTO

	err := d.Database.Select(&users, "select username, userpassword, useruuid from users where useruuid = $1", string(uuid))
	if err != nil {
		//log.Println("In dal.GetUserByUsername", err)
		return domain.UserData{}, err
	}
	if users == nil {
		return domain.UserData{}, service.ErrorUserNotRegistered
	}
	if len(users) == 1 {
		user := userDBDTOToUserData(users[0])
		return user, nil
	}
	if len(users) > 1 {
		return domain.UserData{}, ErrorMoreThanOneUsernameRecord
	}
	return domain.UserData{}, service.ErrorUserNotRegistered
}

func (d *UserDataStoragePGSQL) UpdateUser(uuid domain.UUID, data *domain.UserData) error {
	userDTO := userDataToUserDBDTO(*data)
	userDTO.UUID = string(uuid)

	_, err := d.Database.NamedExec(`UPDATE users SET username=:username, userpassword=:userpassword WHERE useruuid=:useruuid`,
		userDTO)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}
	return nil
}

func (d *UserDataStoragePGSQL) RemoveUser(uuid domain.UUID) error {
	_, err := d.Database.Exec("DELETE FROM users WHERE useruuid=$1", uuid)
	if err != nil {
		//log.Println("In dal.AddNewUser", err)
		return err
	}
	return nil
}

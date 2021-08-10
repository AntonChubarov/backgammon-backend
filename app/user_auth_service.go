package app

import (
	"backgammon/domain"
	"log"
)

type UserAuthService struct {
	storage domain.UserStorage
}

func NewUserAuthService(storage domain.UserStorage) *UserAuthService {
	return &UserAuthService{storage: storage}
}

func (uas *UserAuthService) RegisterNewUser(data domain.UserData) error {
	userExist, err := uas.storage.IsUserExist(data.Login)
	if userExist {
		return domain.UserExistError
	}

	data.UUID = GenerateUUID()

	data.Password, err = HashPassword(data.Password)
	if err != nil {
		log.Println("In app.RegisterNewUser", err)
		return err
	}

	err = uas.storage.AddNewUser(data)
	if err != nil {
		log.Println("In app.RegisterNewUser", err)
		return err
	}

	return nil
}

func (uas *UserAuthService) AuthorizeUser(data domain.UserData) (token string, err error) {
	token = ""
	var userExist bool
	var user domain.UserData

	userExist, err = uas.storage.IsUserExist(data.Login)
	if !userExist {
		err = domain.InvalidLogin
		return
	}

	user, err = uas.storage.GetUserByLogin(data.Login)
	if err != nil {
		return
	}
	var passwordHash string
	passwordHash, err = HashPassword(data.Password)
	if err != nil {
		return
	}

	if passwordHash != user.Password {
		err = domain.InvalidPassword
		return
	}

	token = "Some token from service"
	return
}
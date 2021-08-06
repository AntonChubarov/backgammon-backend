package app

import (
	"backgammon/domain"
	"log"
)

type UserRegistrationService struct {
	storage domain.UserStorage
}

func NewUserRegistrationService(storage domain.UserStorage) *UserRegistrationService {
	return &UserRegistrationService{storage: storage}
}

func (urs *UserRegistrationService) RegisterNewUser(data domain.UserData) error {
	userExist, err := urs.storage.IsUserExist(data)
	if userExist {
		return domain.UserExistError
	}

	data.UUID = GenerateUUID()

	data.Password, err = HashPassword(data.Password)
	if err != nil {
		log.Println("In app.RegisterNewUser", err)
		return err
	}

	err = urs.storage.AddNewUser(data)
	if err != nil {
		log.Println("In app.RegisterNewUser", err)
		return err
	}

	return nil
}